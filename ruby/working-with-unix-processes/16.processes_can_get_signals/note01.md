# [Processes Can Get Signals](https://workingwithruby.com/wwup/signals/)

+ TODO read this part again

Confusing

+ In the last chapter we looked at `Process.wait`. It provides a nice way for a parent process to keep tabs on its child processes. However Not every parent has the luxury of waiting around on their children all day.

上一节里讲了`Process.wait`方法; 它提供了从父进程里关注fork出的子进程的方式;

然而`Process.wait`是一个阻塞调用, 在子进程死亡前, 他不会返回

+ What’s a busy parent to do? Not every parent has the luxury of waiting around on their children all day.

如果父进程很忙碌怎么办(不能接受阻塞怎么办)? 不是所有父进程都有时间等子进程返回的

+ There is a solution for the busy parent! And it’s our introduction to Unix signals.

对于忙碌的父进程, 这里有一种解决办法: unix 信号

## Trapping SIGCHLD

把上一节的例子改造为"一个忙碌的父进程"的例子

+ man sigaction
    + SIGCHLD: child status has changed

+ 这个版本**有时候**没法正常退出, 因为信号量不可靠; 改为后面用循环调用`Process.wait`的形式就没问题了
```ruby
child_processes = 3
dead_processes = 0
# We fork 3 child processes.
child_processes.times do
  fork do
    # They sleep for 3 seconds.
    sleep 3
  end
end

# Our parent process will be busy doing some intense mathematics.
# But still wants to know when one of its children exits.

# By trapping the :CHLD signal our process will be notified by the kernel
# when one of its children exits.
trap(:CHLD) do
  # Since Process.wait queues up any data that it has for us we can ask for it
  # here, since we know that one of our child processes has exited.

  puts Process.wait
  dead_processes += 1
  # We exit explicitly once all the child processes are accounted for.
  exit if dead_processes == child_processes
end

# Work it.
loop do
  (Math.sqrt(rand(44)) ** 8).floor
  sleep 1
end
```

## SIGCHLD and Concurrency

+ Signal delivery is unreliable.
    + By this I mean that if your code is handling a CHLD signal while another child process dies you may or may not receive a second CHLD signal.

信号的派发是不可靠的

即 如果你的代码trap了信号量, 同时有一个子进程死亡了, 这时候你的程序有可能收到`CHLD`信号, 也可能收不到这个信号

+ To properly handle CHLD you must call `Process.wait` in a loop and look for as many dead child processes as are available, since you may have received multiple CHLD signals since entering the signal handler.

想要处理好`CHLD`信号, 我们循环调用`Process.wait`, 以确保已经死亡的子进程资源被回收

但是`Process.wait`是一个阻塞方法, 这种阻塞可以通过给它传入第二个参数来处理

```ruby
Process.wait(-1, Process::WNOHANG)
```

`Process::WNOHANG`告诉内核: 如果没有子进程退出, 不需要阻塞

+ Here’s a rewrite of the code snippet from the beginning of this chapter that won’t ‘miss’ any child process deaths:

+ TODO unclear about `$stdout.sync = true`

```ruby
child_processes = 3
dead_processes = 0
# We fork 3 child processes.
child_processes.times do
  fork do
    # They sleep for 3 seconds.
    sleep 3
  end
end

# Sync $stdout so the call to #puts in the CHLD handler isn't
# buffered. Can cause a ThreadError if a signal handler is
# interrupted after calling #puts. Always a good idea to do
# this if your handlers will be doing IO.
$stdout.sync = true

# Our parent process will be busy doing some intense mathematics.
# But still wants to know when one of its children exits.

# By trapping the :CHLD signal our process will be notified by the kernel
# when one of its children exits.
trap(:CHLD) do
  # Since Process.wait queues up any data that it has for us we can ask for it
  # here, since we know that one of our child processes has exited.

  # We loop over a **non-blocking Process.wait** to ensure that any dead child
  # processes are accounted for.
  begin
    while pid = Process.wait(-1, Process::WNOHANG) # -1 means Waits for any child process (the default if no pid is given).
      puts pid
      dead_processes += 1
    end
  rescue Errno::ECHILD
  end
end

loop do
  # We exit ourself once all the child processes are accounted for.
  exit if dead_processes == child_processes

  sleep 1
end
```

## Signals Primer

+ Signals are asynchronous communication.

+ When a process receives a signal from the kernel it can do one of the following:
    + ignore the signal
    + perform a specified action
    + perform the default action

+ Technically signals are sent by the kernel
    + Signals are sent from one process to another process, using the kernel as a middleman.
    + The original purpose of signals was to specify different ways that a process should be killed.

+ Below is a table showing signals commonly supported on Unix systems.

Signal    | Value      |  Action  | Comment
----------|------------|----------|----------------------------------------------
SIGHUP    |     1      | Term     |  Hangup detected on controlling terminal or death of controlling process
SIGINT    |     2      | Term     |  Interrupt from keyboard
SIGQUIT   |     3      | Core     |  Quit from keyboard
SIGILL    |     4      | Core     |  Illegal Instruction
SIGABRT   |     6      | Core     |  Abort signal from abort(3)
SIGFPE    |     8      | Core     |  Floating point exception
SIGKILL   |     9      | Term     |  Kill signal
SIGSEGV   |    11      | Core     |  Invalid memory reference
SIGPIPE   |    13      | Term     |  Broken pipe: write to pipe with no readers
SIGALRM   |    14      | Term     |  Timer signal from alarm(2)
SIGTERM   |    15      | Term     |  Termination signal
SIGUSR1   | 30,10,16   | Term     |  User-defined signal 1
SIGUSR2   | 31,12,17   | Term     |  User-defined signal 2
SIGCHLD   | 20,17,18   | Ign      |  Child stopped or terminated
SIGCONT   | 19,18,25   | Cont     |  Continue if stopped
SIGSTOP   | 17,19,23   | Stop     |  Stop process
SIGTSTP   | 18,20,24   | Stop     |  Stop typed at tty
SIGTTIN   | 21,21,26   | Stop     |  tty input for background process
SIGTTOU   | 22,22,27   | Stop     |  tty output for background process


The signals SIGKILL and SIGSTOP cannot be trapped, blocked, or ignored.

+ It’s interesting to note the SIGUSR1 and SIGUSR2 signals.
    + These are signals whose action is meant specifically to be defined by your process.
    + We’ll see shortly that we’re free to redefine any of the signal actions that we please, but those two signals are meant for your use.

+ Redefining Signals
```ruby
puts Process.pid
trap(:INT) { print "Na na na, you can't get me" }
sleep # so that we have time to send it a signal

Process.kill(:INT, <pid of first session>)
# ctrl+c won't work either

Process.kill(:TERM, <pid of first session>)
Process.kill(:KILL, <pid of first session>)
```

+ Ignoring Signals
```ruby
puts Process.pid
trap(:INT, "IGNORE")
sleep # so th

Process.kill(:INT, <pid of first session>)
```

+ It’s good to keep in mind that trapping a signal is a bit **like using a global variable**, you might be overwriting something that some other code depends on. And unlike global variables signal handlers can’t be namespaced.

+ There is a way to preserve handlers defined by other Ruby code, so that your signal handler won’t trample any other ones that are already defined.
```ruby
trap(:INT) { puts 'This is the first signal handler' }

old_handler = trap(:INT) {
  old_handler.call
  puts 'This is the second handler'
  exit
}
sleep 5 # so that we have time to send it a signal
```

```ruby
system_handler = trap(:INT) {
  puts 'about to exit!'
  system_handler.call
}
sleep 5 # so that we have time to send it a signal
```

+ In terms of **best practices** your code probably shouldn't define any signal handlers, unless it's a server.
+ It's very rare that library code should trap a signal.

```ruby
# The 'friendly' method of trapping a signal.

old_handler = trap(:QUIT) {
  # do some cleanup
  puts 'All done!'

  old_handler.call if old_handler.respond_to?(:call)
}
```

+ This violates the expectations of ......

+ Whether or not you decide to preserve previously defined signal handlers is up to you, **just make sure you know why you’re doing it**.

+ Your process can receive a signal anytime. That’s the beauty of them! They’re asynchronous.

+ With signals, any process can communicate with any other process on the system, so long as it knows its pid.

+ This makes signals a very powerful communication tool.

+ It’s common to send signals from the shell using kill(1).

+ And for the most part it will be the human users who are sending signals rather than automated programs.

+ Ruby’s Process.kill maps to kill(2), Kernel#trap maps roughly to sigaction(2). signal(7) is also useful.



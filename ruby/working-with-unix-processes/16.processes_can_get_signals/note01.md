https://workingwithruby.com/wwup/signals/

+ TODO read this part again

Confusing

+ In the last chapter we looked at Process.wait. It provides a nice way for a parent process to keep tabs on its child processes. However Not every parent has the luxury of waiting around on their children all day.

+ man sigaction
    + SIGCHLD

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

+ Signal delivery is unreliable.
    + By this I mean that if your code is handling a CHLD signal while another child process dies you may or may not receive a second CHLD signal.

+ Here’s a rewrite of the code snippet from the beginning of this chapter that won’t ‘miss’ any child process deaths:

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
    while pid = Process.wait(-1, Process::WNOHANG)
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




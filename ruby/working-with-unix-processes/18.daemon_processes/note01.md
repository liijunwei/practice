# [Daemon Processes](https://workingwithruby.com/wwup/daemons/)

+ Important

+ TODO read again and again

## The First Process

+ There is one daemon process in particular that has special significance for your operating system.
    + What is the pid of the very first process on the system?
    + This is a classic who-created-the-creator kind of problem, and it has a simple answer.
    + When the kernel is bootstrapped it spawns a process called **the init process**
    + This process has a ppid of 0 and is the ‘grandparent of all processes’.
    + It’s the first one and it has no ancestor. Its pid is 1.

## Creating Your First Daemon Process

+ Any process can be made into a daemon process.

+ `rack project` as example

## [Diving into Rack](https://github.com/rack/rack/commit/485c92c4adc3dc48491e2a9fd742c01584e36760)

+ rack demo

```ruby
def daemonize_app
  if RUBY_VERSION < "1.9"
    exit if fork # makes use of the return value of the fork method
                 # As always, the return value will be truth-y for the parent and false-y for the child
                 # This means that the parent process will exit, and as we know, orphaned child processes carry on as normal.
    Process.setsid
    exit if fork
    Dir.chdir "/"
    STDIN.reopen "/dev/null"
    STDOUT.reopen "/dev/null", "a"
    STDERR.reopen "/dev/null", "a"
  else
    Process.daemon
  end
end
```

+ [MRI source for Process.daemon](https://github.com/ruby/ruby/blob/c852d76f46a68e28200f0c3f68c8c67879e79c86/process.c#L4817-4860)

## Daemonizing a Process, Step by Step

```ruby
exit if fork
```

+ The ppid of orphaned processes is always 1. This is the only process that the kernel can be sure is active at all times.

+ Calling `Process.setsid` does three things:(in the forked process)
    1. The process becomes a session leader of a new session
    2. The process becomes the process group leader of a new process group
    3. The process has no controlling terminal

## Process Groups and Session Groups

+ Process groups and session groups are all about job control.
+ By "job control" I'm referring to the way that processes are handled by the terminal.

```ruby
puts Process.getpgrp
puts Process.pid
```

+ Each and every process belongs to a group, and each group has a unique integer id.

每个进程都属于一个进程分组, 每个分组都有一个唯一的id

+ A process group is just a collection of related processes

一个进程分组就是一组相关的进程

+ Typically the process group id will be the same as the pid of the process group leader.

一般来说, 进程分组的id和这个进程分组的leader(unclear ???)相同

+ The process group leader is the "originating" process of a terminal command.
    + ie. If you start an irb process at the terminal it will become the group leader of a new process group.
    + Any child processes that it creates will be made part of the same process group.

今晨分组的leader是终端命令的"发起"进程

例如, 我们在命令行里启动了irb, 然后在irb进程里fork出新的进程, 那么irb就是这个它新发起的进程的进程组leader

+ Typically the **process group** id will be the same as the pid of the process group leader.

+ Try out the following example to see that process groups are inherited.
```ruby
puts Process.pid
puts Process.getpgrp

fork {
  puts Process.pid
  puts Process.getpgrp
}
```

+ Child processes are not given special treatment by the kernel.
+ Exit a parent process and the child will continue on.
+ This is the behaviour when a parent process exits, but the behaviour is a bit different when the parent process is being controlled by a terminal and is killed by a signal.

+ The terminal receives the signal and forwards it on to any process in the foreground process group.
+ In this case, both the Ruby script and the long-running shell command would part of the same process group, so they would both be killed by the same signal.

+ A **session group** is one level of abstraction higher up, a collection of process groups.
```bash
git log | grep shipped | less

# Even though these commands are not part of the same process group one Ctrl-C will kill them all.
# These commands are part of the same session group.
# Each invocation from the shell gets its own session group.
# An invocation may be a single command or a string of commands joined by pipes.
```

+ Again, your terminal handles session groups in a special way:
    + sending a signal to the session leader will forward that signal to all the process groups in that session
    + which will forward it to all the processes in those process groups

+ we want to be fully detached from a terminal.
    + `Process.setsid` will make this forked process the leader of a new process group and a new session group.
    + Note that `Process.setsid` will fail in a process that is already a process group leader, **it can only be run from child processes.**

+ If you think you want to create a daemon process you should ask yourself one basic question:
    + Does this process need to stay responsive forever?
    + If the answer is no,  then you probably want to look at a cron job or background job system.
    + If the answer is yes, then you probably have a good candidate for a daemon process.

+ system calls
    + Process.setsid  -> setsid(2)
    + Process.getpgrp -> getpgrp(2)





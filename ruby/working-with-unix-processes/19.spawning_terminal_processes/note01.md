# [Spawning Terminal Processes](https://workingwithruby.com/wwup/spawning/)

+ A common interaction in a Ruby program is ["shelling out"](https://stackoverflow.com/questions/28628985/what-does-shell-out-or-shelling-out-mean) from your program to run a command in a terminal.

## fork + exec

+ All of the methods described below are variations on one theme: `fork(2)` + `execve(2)`.

下面描述的方法都是 `fork+ececve` 的变体

+ `execve(2)` allows you to replace the current process with a different process.
    + `man 2 execve`: execve() transforms the calling process into a new process.

+ `execve(2)` allows you to transform the current process into any other process.
    + You can take a Ruby process and turn it into a Python process, or an ls(1) process, or another Ruby process.

+ `execve(2)` transforms the process and never returns.

+ The fork + exec combo is a common one when spawning new processes.

+ `execve(2)` is a very powerful and efficient way to transform the current process into another one; **the only catch is that your current process is gone**. That's where `fork(2)` comes in handy.

+ You can use `fork(2)` to create a new process, then use `execve(2)` to transform that process into anything you like. Your current process is still running just as it was before and you were able to spawn any other process that you want to.

你可以用`fork(2)`创建出新的进程, 然后用`execve(2)`把新建的进程转换为你想要的进程

+ If your program depends on the output from the `execve(2)` call you can use the tools you learned in previous chapters to handle that.
    + `Process.wait` will ensure that your program waits for the child process to finish whatever it's doing so you can get the result back.

## File descriptors and exec

+ At the OS level, a call to `execve(2)` doesn't close any open file descriptors by default.

操作系统层的`execve`调用, 默认不会关闭文件描述符

+ However, a call to `exec` in Ruby will close all open file descriptors by default (excluding the standard streams).

但是在ruby里的`exec`调用默认会关闭标准流以外的文件描述符

+ This default behaviour of closing file descriptors on exec prevents file descriptor 'leaks'.

+ In this example we start up a Ruby program and open the `/etc/hosts` file.
    + Then we `exec` a `python` process and tell it to open the file descriptor number that Ruby received for opening the `/etc/hosts` file.
    + You can see that python recognizes this file descriptor (because it was shared via `execve(2)`) and is able to read from it without having to open the file again.
```ruby
hosts = File.open('/etc/hosts')

python_code = %Q[import os; print os.fdopen(#{hosts.fileno}).read()]

# The hash as the last arguments maps any file descriptors that should
# stay open through the exec.
exec 'python', '-c', python_code, {hosts.fileno => hosts}
```

+ Unlike `fork(2)`, `execve(2)` does not share memory with the newly created process.

和fork不同, execve不和新建的进程共享内存

## Arguments to exec

+ Pass a string to exec and it will actually start up a shell process and pass the string to the shell to interpret.

+ Pass an array and it will skip the shell and set up the array directly as the ARGV to the new process.

+ **Generally you want to avoid passing a string unless you really need to. Pass an array where possible.**

+ Passing a string and running code through the shell can raise security concerns.

直接用字符串shell out, 会有安全风险

## Kernel#system

```ruby
system('ls')
system('ls', '--help')
system('git log | tail -10')
```

## "Kernel#\`"
```bash
The return value of Kernel#system reflects the exit code of the terminal command in the most basic way.

Kernel#` works slightly differently. The value returned is the STDOUT of the terminal program collected into a String.

Kernel#` and %x[] do the exact same thing.
```

## Process.spawn

```ruby
# This call will start up the 'rails server' process with the
# RAILS_ENV environment variable set to 'test'.
Process.spawn({'RAILS_ENV' => 'test'}, 'rails server')

# This call will merge STDERR with STDOUT for the duration
# of the 'ls --help' program.
Process.spawn('ls', '--zz', STDERR => STDOUT)
```

**TODO** 2022-05-16 09:30:51

+ `Process.spawn` is a bit different than the others in that **it is non-blocking**.
```ruby
# Do it the blocking way
system 'sleep 5'

# Do it the non-blocking way
Process.spawn 'sleep 5'

# Do it the blocking way with Process.spawn
# Notice that it returns the pid of the child process
pid = Process.spawn 'sleep 5'
Process.waitpid(pid)
```

+ The last example in this code block is a really great example of the flexibility of Unix programming.

+ In previous chapters we talked a lot about `Process.wait`, but it was always in the context of forking and then running some Ruby code.

+ You can see from this example that the kernel cares not what you are doing in your process, it will always work the same.

+ **consistency**

+ All code looks the same to the kernel; that's what makes it such a flexible system.

这些代码在内核看来没有什么区别, 正是这种抽象使得系统很灵活

+ [Process.spawn](http://www.ruby-doc.org/core-1.9.3/Process.html#method-c-spawn) takes many options that allow you to control the behaviour of the child process.


## IO.popen

+ [gem: whenever](https://github.com/javan/whenever) usees this to write crontab

```ruby
# This example will return a file descriptor (IO object). Reading from it
# will return what was printed to STDOUT from the shell command.
IO.popen('ls')
```

+ The most common usage for `IO.popen` is an implementation of Unix pipes in pure Ruby.

`IO.popen`最常见的用法就是 使用ruby实现unix的管道

+ That's where the 'p' comes from in popen.

这就是`popen`里"p"的来源

+ Underneath it's still doing the `fork+exec`, but it's also setting up a pipe to communicate with the spawned process.

`IO.popen`背后还是使用`fork+exec`, 不过它还结合了管道用以和其他进程沟通

+ That pipe is passed as the block argument in the block form of IO.popen.

```ruby
# An IO object is passed into the block. In this case we open the stream
# for writing, so the stream is set to the STDIN of the spawned process.
#
# If we open the stream for reading (the default) then
# the stream is set to the STDOUT of the spawned process.
IO.popen('less', 'w') { |stream|
  stream.puts "some\ndata"
}
```

## Open3

+ `Open3` allows simultaneous access to the STDIN, STDOUT, and STDERR of a spawned process.
+ Open3 acts like a more flexible version of IO.popen, for those times when you need it.
```ruby
# This is available as part of the standard library.
require 'open3'

Open3.popen3('grep', 'data') { |stdin, stdout, stderr|
  stdin.puts "some\ndata"
  stdin.close
  puts stdout.read
}

# Open3 will use Process.spawn when available. Options can be passed to
# Process.spawn like so:
Open3.popen3('ls', '-uhh', :err => :out) { |stdin, stdout, stderr|
  puts stdout.read
}
```

## In the Real World

+ One drawback to all of these methods is that they rely on `fork(2)`.
    + What's wrong with that? Imagine this scenario:
    + You have a big Ruby app that is using hundreds of MB of memory.
    + You need to shell out. If you use any of the methods above you'll incur the cost of forking.

有一点需要注意: 上面提到的这些ruby的功能, 都用到了`fork`, 所以, 在shell out的时候, 内存使用率可能会出问题

+ When you `fork(2)` the process the kernel doesn't know that you're about to transform that process with an exec(2). You may be forking in order to run Ruby code, in which case you'll need to have all of the memory available.

+ It's good to keep in mind that `fork(2)` has a cost, and sometimes it can be a performance bottleneck.
+ Question: What if you need to shell out a lot and don't want to incur the cost of `fork(2)`?
    + There are some native Unix system calls for spawning processes without the overhead of `fork(2)`.
    + Unfortunately they don't have support in the Ruby language core library.
    + However, there is a Rubygem that provides a Ruby interface to these system calls.
    + The [posix-spawn](https://github.com/rtomayko/posix-spawn/) project provides access to posix_spawn(2), which is available on most Unix systems.

+ At a basic level posix_spawn(2) is a subset of `fork(2)`.

+ Recall the two discerning(挑剔的) attributes of a new child process from `fork(2)`:
    1. it gets an exact copy of everything that the parent process had in memory
    2. it gets a copy of all the file descriptors that the parent process had open.
    + posix_spawn(2) preserves #2, but not #1. That's the big difference between the two.

+ system calls
    + Kernel#system maps to system(3)
    + Kernel#exec   maps to `execve(2)`
    + IO.popen      maps to popen(3)
    + posix-spawn   uses    posix_spawn(2)


# [Spawning Terminal Processes](https://workingwithruby.com/wwup/spawning/)

+ A common interaction in a Ruby program is [‘shelling out’](https://stackoverflow.com/questions/28628985/what-does-shell-out-or-shelling-out-mean) from your program to run a command in a terminal.

+ execve(2) allows you to replace the current process with a different process.
    + `man 2 execve`: execve() transforms the calling process into a new process.

+ execve(2) allows you to transform the current process into any other process.
    + You can take a Ruby process and turn it into a Python process, or an ls(1) process, or another Ruby process.

+ execve(2) transforms the process and never returns.

+ The fork + exec combo is a common one when spawning new processes.

+ execve(2) is a very powerful and efficient way to transform the current process into another one; the only catch is that your current process is gone. That’s where fork(2) comes in handy.

+ You can use fork(2) to create a new process, then use execve(2) to transform that process into anything you like. Your current process is still running just as it was before and you were able to spawn any other process that you want to.

+ If your program depends on the output from the execve(2) call you can use the tools you learned in previous chapters to handle that.
    + `Process.wait` will ensure that your program waits for the child process to finish whatever it’s doing so you can get the result back.

+ At the OS level, a call to execve(2) doesn’t close any open file descriptors by default.
    + However, a call to exec in Ruby will close all open file descriptors by default (excluding the standard streams).
    + This default behaviour of closing file descriptors on exec prevents file descriptor ‘leaks’.

+ In this example we start up a Ruby program and open the /etc/hosts file.
    + Then we exec a python process and tell it to open the file descriptor number that Ruby received for opening the /etc/hosts file.
    + You can see that python recognizes this file descriptor (because it was shared via execve(2)) and is able to read from it without having to open the file again.
```ruby
hosts = File.open('/etc/hosts')

python_code = %Q[import os; print os.fdopen(#{hosts.fileno}).read()]

# The hash as the last arguments maps any file descriptors that should
# stay open through the exec.
exec 'python', '-c', python_code, {hosts.fileno => hosts}
```

+ Unlike fork(2), execve(2) does not share memory with the newly created process.

+ Arguments to exec
    + Pass a string to exec and it will actually start up a shell process and pass the string to the shell to interpret.
    + Pass an array and it will skip the shell and set up the array directly as the ARGV to the new process.
    + **Generally you want to avoid passing a string unless you really need to. Pass an array where possible.**

```bash
The return value of Kernel#system reflects the exit code of the terminal command in the most basic way.

Kernel#` works slightly differently. The value returned is the STDOUT of the terminal program collected into a String.

Kernel#` and %x[] do the exact same thing.
```

+ Process.spawn
```ruby
# This call will start up the 'rails server' process with the
# RAILS_ENV environment variable set to 'test'.
Process.spawn({'RAILS_ENV' => 'test'}, 'rails server')

# This call will merge STDERR with STDOUT for the duration
# of the 'ls --help' program.
Process.spawn('ls', '--zz', STDERR => STDOUT)
```






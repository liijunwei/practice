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


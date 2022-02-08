+ pids             represent running processes
+ file descriptors represent open files

+ Any time that you open a resource in a running process it is assigned a file descriptor number.

+ Any resource that your process opens gets a unique number identifying it.

+ It’s important to note that file descriptors keep track of open resources only. Closed resources are not given a file descriptor number.

+ You may have noticed that when we open a file and ask for its file descriptor number the lowest value we get is 3. What happened to 0, 1, and 2?
    + [Standard Streams](https://workingwithruby.com/wwup/fds/#standard-streams)
    + stdin  0
    + stdout 1
    + stderr 2

+ File descriptors are at the **core** of network programming using sockets, pipes, etc. and are also at the core of any file system operations.

+ **Many methods on Ruby’s IO class map to system calls of the same name.** These include open(2), close(2), read(2), write(2), pipe(2), fsync(2), stat(2), among others.

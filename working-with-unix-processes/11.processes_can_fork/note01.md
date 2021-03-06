https://workingwithruby.com/wwup/forking/

+ Forking is one of the most powerful concepts in Unix programming.

+ `man 2 fork`
+ The fork(2) system call allows a running process to create new process programmatically.
    + This new process is an **exact copy of the original process except for the following**(**go read manpage**):

+ When forking, the process that initiates the fork(2) is called the `parent`, and the newly created process is called the `child`.

+ **The child process inherits a copy of all of the memory in use by the parent process, as well as any open file descriptors belonging to the parent process.**

子进程从父进程的内存里继承了一个完全的副本, 包括已打开的文件的文件描述符

+ The child process inherits a copy of everything that the parent process has in main memory. In this way a process could load up a large codebase

子进程从父进程复制了所有的东西, 从父进程fork出的进程能加载到很大的代码库

+ Say a Rails app, that occupies 500MB of main memory. Then this process can fork 2 new child processes. Each of these child processes would effectively have their own copy of that codebase loaded in memory.

例如一个rails应用占用了500M内存, 然后他fork出两个子进程; 每个子进程的内存里都加载了相同的代码

+ The call to fork returns near-instantly so we now have 3 processes with each using 500MB of memory

调用fork产生两个新的进程几乎不花什么时间, 然后你就能获得3个一样的进程啦

+ Perfect for when you want to have multiple instances of your application loaded in memory at the same time. Because only one process needs to load the app and forking is fast, this method is faster than loading the app 3 times in separate instances.

当你想要多个进程的时候, 使用先启动一个,再让他fork出两个一样的子进程很管用, 因为fork很快; 它比同时启动三个进程要快得多

+ The child process inherits any open file descriptors from the parent at the time of the fork(2). It’s given the same map of file descriptor numbers that the parent process has. In this way the two processes can share open files, sockets, etc.

+ See the next chapter for a discussion of **copy-on-write**(TODO) and how it affects memory when forking.

```ruby
if fork
  puts "entered the if block"
else
  puts "entered the else block"
end
```

+ It’s no mystery what’s happening here
    + **One call to the fork method actually returns twice** 调用一次fork方法, 会有两次返回: 一次在父进程, 一次在子进程
    + Remember that fork creates a new process
    + So it returns once in the calling process (parent) and once in the newly created process (child)
    + In the child process fork returns nil. Since nil is falsy it executes the code in the else block.
    + In the parent process fork returns the pid of the newly created child process. Since an integer is truthy it executes the code in the if block.

```ruby
puts "parent process pid is #{Process.pid}"

# In the parent process fork returns the pid of the newly created child process
if pid = fork
  puts "entered the if block from #{Process.pid}, new process pid is #{pid}"
else
  # executed by the child process
  # In the child process fork returns nil
  puts "entered the else block from #{Process.pid}"
end
```

+ This concept is illustrated nicely by simply printing the return value of a fork call.
```ruby
# in parent process: fork returns pid of child process
# in child process:  fork returns nil
puts fork
```

+ fork(2) creates a new process that's a copy of the old process.

+ It’s also possible, and more common in Ruby code, to use fork with a block

在ruby中, 调用fork创建子进程时, 跟常见的方法是和块结合使用

+ When you pass a block to the fork method that block will be executed in the new child process, while the parent process simply skips over it.

当fork和块一起使用时, 块里的代码在父进程中会被跳过

+ The child process exits when it’s done executing the block. It does not continue along the same code path as the parent.

块里的代码会在子进程里执行, 并且执行完毕后会退出

+ Ruby’s Kernel#fork maps to fork(2).

+ man 2 fork
```
RETURN VALUES
     Upon successful completion, fork() returns a value of 0 to the child process and returns the process ID of the child process to the parent process.
     Otherwise, a value of -1 is returned to the parent process, no child process is created, and the global variable errno is set to indicate the error.
```







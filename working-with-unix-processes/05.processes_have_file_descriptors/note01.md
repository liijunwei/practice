+ pids             represent running processes
+ file descriptors represent open files

+ part of "the Unix philosophy"
    + "everything is a file"
    + devices/sockets/pipes ... -> resources
    + files                     -> classical files

+ Any time that you open a resource in a running process it is assigned a file descriptor number.

进程打开一个资源后, 系统会给这个资源分配一个文件描述符号码(file descriptor number)

+ File descriptors are NOT shared between unrelated processes, they live and die with the process they are bound to, just as any open resources for a process are closed when it exits.

文件描述符不在无关的进程间共享, 他们只和当前进程相关

+ There are special semantics for file descriptor sharing when you fork a process, more on that later.

在fork进程时, 文件描述符有特殊的语义, 之后详解

> In Ruby, open resources are represented by the `IO` class. Any `IO` object can have an associated file descriptor number. Use `IO#fileno` to get access to it.

在ruby中, 打开的资源由IO类的实例表示

每个IO实例都有对应的文件描述符, 可以由api `IO#fileno` 查看

```ruby
passwd = File.open('/etc/passwd')
puts passwd.fileno
```

+ Any resource that your process opens gets a unique number identifying it.

进程打开的所有资源, 都会有一个对应的文件描述符去标识它

+ This is how the kernel keeps track of any resources that your process is using.

这就是内核跟踪进程正在使用的资源的方法

+ It’s important to note that file descriptors keep track of open resources only. Closed resources are not given a file descriptor number.

需要注意: 文件描述符只追踪打开的资源, 操作系统不会为关闭的资源分配文件描述符

+ Given the above, file descriptors are sometimes called "open file descriptors".

根据前面的描述, 文件描述符也可以叫做 "打开的文件的描述符"

+ You may have noticed that when we open a file and ask for its file descriptor number the lowest value we get is 3. What happened to 0, 1, and 2?

你可能注意到了, 当我们打开文件后, 文件描述符的号码是从 **3** 开始的, 那么 0, 1, 2 去哪里了?

+ Every Unix process comes with three open resources: [Standard Streams](https://workingwithruby.com/wwup/fds/#standard-streams)
    + stdin  0(STDIN):  provides a generic way to read input from keyboard devices or pipes
    + stdout 1(STDOUT): generic ways to write standard output to monitors, files, printers, etc
    + stderr 2(STDERR): generic ways to write error    output to monitors, files, printers, etc

每个unix进程都至少有3个打开的资源: STDIN/STDOUT/STDERR

STDIN提供了从键盘设备或管道读取输入的通用方法
STDOUT和STDERR提供了将输出写入到显示器/文件/打印机等的通用方法

```ruby
puts STDIN.fileno
puts STDOUT.fileno
puts STDERR.fileno
```

+ File descriptors are at the **core** of network programming using sockets, pipes, etc. and are also at the core of any file system operations.

文件描述符是网络编程中使用 socket/pipe 等的核心

文件描述符也是任何文件系统操作的核心

+ **Many methods on Ruby’s IO class map to system calls of the same name.** These include open(2), close(2), read(2), write(2), pipe(2), fsync(2), stat(2), among others.

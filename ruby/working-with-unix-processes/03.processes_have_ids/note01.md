+ Every process running on your system has a unique process identifier, hereby referred to as ‘pid’.

所有运行在系统上的进程都有一个唯一的标识符: 进程ID(PID)

+ This is how the kernel sees your process: as a number.

在系统内核看来, 进程就是用PID区分的

```ruby
puts Process.pid
```

> A pid is a simple, generic representation of a process.

+ scenario for PID
    + multiple processes logging to same file, and use pid to identify each other

+ System Calls
    + `Process.pid` maps to `man 2 getpid`

+ A ruby global variable holds value of current pid: `$$`
    + use `Process.pid` when possible and avoid `$$`


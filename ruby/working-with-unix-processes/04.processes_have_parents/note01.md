+ Every process running on your system has a parent process.

系统里的每个进程都有一个父进程

+ Each process knows its parent process identifier (hereby referred to as ‘ppid’).

每个进程都知道它的父进程的唯一标识: PPID

```ruby
Process.pid  # to get        pid of current process
Process.ppid # to get parent pid of current process
```

+ Ruby’s `Process.ppid` maps to `man 2 getppid`


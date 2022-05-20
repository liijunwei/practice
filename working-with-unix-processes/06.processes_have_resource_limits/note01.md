https://workingwithruby.com/wwup/rlimits/

+ fact: open resources are represented by file descriptors

进程使用中的资源是用文件描述符表示的

+ Q: how many file descriptors can one process have?
    + The answer depends on your system configuration, but the important point is there are some resource limits imposed on a process by the kernel.

每个进程能打开的文件描述符的数量是有限制的, 限制来自于内核

```ruby
# getrlimit(resource) → [cur_limit, max_limit]
Process.getrlimit(:NOFILE)
=> [256, 9223372036854775807] # Darwin bxzy 20.6.0 Darwin Kernel Version 20.6.0: Wed Nov 10 22:23:07 PST 2021; root:xnu-7195.141.14~1/RELEASE_X86_64 x86_64
=> [65535, 65535] # Linux xiaoli 4.15.0-117-generic #118-Ubuntu SMP Fri Sep 4 20:02:41 UTC 2020 x86_64 x86_64 x86_64 GNU/Linux
```

+ Soft Limits vs. Hard Limits#
    + If you’re interested in changing the limits at a system-wide level then start by having a look at sysctl(8).
    + `man 8 sysctl`


+ Note that exceeding the soft limit will raise `Errno::EMFILE`

```ruby
# The maximum number of simultaneous processes
# allowed for the current user.
Process.getrlimit(:NPROC)

# The largest size file that may be created.
Process.getrlimit(:FSIZE)

# The maximum size of the stack segment of the
# process.
Process.getrlimit(:STACK)
```

+ Ruby’s `Process.getrlimit` and `Process.setrlimit` map to `getrlimit(2)` and `setrlimit(2)`, respectively.


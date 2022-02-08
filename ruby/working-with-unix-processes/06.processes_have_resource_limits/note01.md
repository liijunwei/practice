https://workingwithruby.com/wwup/rlimits/

+ Q: how many file descriptors can one process have?
    + The answer depends on your system configuration, but the important point is there are some resource limits imposed on a process by the kernel.

```ruby
# getrlimit(resource) → [cur_limit, max_limit]
Process.getrlimit(:NOFILE)
=> [256, 9223372036854775807] # Darwin bxzy 20.6.0 Darwin Kernel Version 20.6.0: Wed Nov 10 22:23:07 PST 2021; root:xnu-7195.141.14~1/RELEASE_X86_64 x86_64
=> [65535, 65535] # Linux xiaoli 4.15.0-117-generic #118-Ubuntu SMP Fri Sep 4 20:02:41 UTC 2020 x86_64 x86_64 x86_64 GNU/Linux
```


# [Daemon Processes](https://workingwithruby.com/wwup/daemons/)

+ There is one daemon process in particular that has special significance for your operating system.
    + What is the pid of the very first process on the system?
    + This is a classic who-created-the-creator kind of problem, and it has a simple answer.
    + When the kernel is bootstrapped it spawns a process called **the init process**
    + This process has a ppid of 0 and is the ‘grandparent of all processes’.
    + It’s the first one and it has no ancestor. Its pid is 1.

+ Any process can be made into a daemon process.

+ rack demo
```ruby
def daemonize_app
  if RUBY_VERSION < "1.9"
    exit if fork
    Process.setsid
    exit if fork
    Dir.chdir "/"
    STDIN.reopen "/dev/null"
    STDOUT.reopen "/dev/null", "a"
    STDERR.reopen "/dev/null", "a"
  else
    Process.daemon
  end
end
```

+ [MRI source for Process.daemon](https://github.com/ruby/ruby/blob/c852d76f46a68e28200f0c3f68c8c67879e79c86/process.c#L4817-4860)






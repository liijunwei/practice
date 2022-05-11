https://workingwithruby.com/wwup/exitcodes/

+ When a process comes to an end it has one last chance to make its mark on the world: its exit code.

+ code(0-255)
    + 0: success
    + other: error/failure

+ It’s usually a good idea to stick with the ‘0 as success’ exit code tradition so that your programs will play nicely with other Unix tools.

遵循 把0当做程序正常退出的约定, 能让你的程序更好的和其他程序协作


## How to Exit a Process

There are several ways you can exit a process in Ruby, each for different purposes.

### Kernel#exit

+ `exit`    # This will exit the program with the success status code (0).
+ `exit 22` # You can pass a custom exit code to this method
+  **When Kernel#exit is invoked, before exiting Ruby invokes any blocks defined by Kernel#at_exit**.

```ruby
at_exit { puts 'Last!' }
exit
```

### Kernel#exit!

+ [Kernel#exit!](https://workingwithruby.com/wwup/exitcodes/#exit-1)
    + default exit code is 1
    + do not invode at_exit block

### Kernel#abort

+ [Kernel#abort](https://workingwithruby.com/wwup/exitcodes/#abort) provides a generic way to exit a process unsuccessfully.
+ Kernel#abort will set the exit code to 1 for the current process.
+ can pass message to stderr
+ Kernel#at_exit blocks are invoked when using Kernel#abort.

### Kernel#raise

+ A different way to end a process is with an unhandled exception.
+ This is something that you never want to happen in a production environment, but it’s almost always happening in development and test environments.
+ Ending a process this way will still invoke any at_exit handlers and will print the exception message and backtrace to STDERR.


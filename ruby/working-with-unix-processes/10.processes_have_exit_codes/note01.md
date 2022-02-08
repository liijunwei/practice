https://workingwithruby.com/wwup/exitcodes/

+ When a process comes to an end it has one last chance to make its mark on the world: its exit code.

+ code(0-255)
    + 0: success
    + other: error/failure

+ There are several ways you can exit a process in Ruby, each for different purposes.
    + Kernel#exit
    + `exit`    # This will exit the program with the success status code (0).
    + `exit 22` # You can pass a custom exit code to this method
    +  **When Kernel#exit is invoked, before exiting Ruby invokes any blocks defined by Kernel#at_exit**.
```ruby
at_exit { puts 'Last!' }
exit
```
    + [Kernel#exit!](https://workingwithruby.com/wwup/exitcodes/#exit-1)
        + default exit code is 1
        + do not invode at_exit block
    + [Kernel#abort](https://workingwithruby.com/wwup/exitcodes/#abort) provides a generic way to exit a process unsuccessfully. Kernel#abort will set the exit code to 1 for the current process.
    + Kernel#raise




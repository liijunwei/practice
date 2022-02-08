https://workingwithruby.com/wwup/forking/

+ Forking is one of the most powerful concepts in Unix programming.

+ `man 2 fork`
+ The fork(2) system call allows a running process to create new process programmatically.
    + This new process is an exact copy of the original process except for the following(**go read manpage**):

+ When forking, the process that initiates the fork(2) is called the `parent`, and the newly created process is called the `child`.

+ **The child process inherits a copy of all of the memory in use by the parent process, as well as any open file descriptors belonging to the parent process.**

+ e.g.
    + The child process inherits a copy of everything that the parent process has in main memory. In this way a process could load up a large codebase
    + Say a Rails app, that occupies 500MB of main memory. Then this process can fork 2 new child processes. Each of these child processes would effectively have their own copy of that codebase loaded in memory.
    + The call to fork returns near-instantly so we now have 3 processes with each using 500MB of memory
    + Perfect for when you want to have multiple instances of your application loaded in memory at the same time. Because only one process needs to load the app and forking is fast, this method is faster than loading the app 3 times in separate instances.

+ See the next chapter for a discussion of **copy-on-write**(TODO) and how it affects memory when forking.

```ruby
if fork
  puts "entered the if block"
else
  puts "entered the else block"
end
```

+ It’s no mystery what’s happening here
    + One call to the fork method actually returns twice
    + Remember that fork creates a new process
    + So it returns once in the calling process (parent) and once in the newly created process (child)
    + In the child process fork returns nil. Since nil is falsy it executes the code in the else block.
    + In the parent process fork returns the pid of the newly created child process. Since an integer is truthy it executes the code in the if block.

```ruby
puts "parent process pid is #{Process.pid}"

if fork
  puts "entered the if block from #{Process.pid}"
else
  puts "entered the else block from #{Process.pid}"
end
```

+ This concept is illustrated nicely by simply printing the return value of a fork call.
```ruby
puts fork
```



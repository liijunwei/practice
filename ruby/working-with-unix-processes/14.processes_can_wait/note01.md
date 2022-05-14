https://workingwithruby.com/wwup/wait/

如果父进程先于fork出的子进程退出了, 可能会出现一些奇怪的现象

+ "Fire a fork and continue in parent process" scenario is really only suitable for one use case, fire and forget.

```ruby
message = 'Good Morning'
recipient = 'tree@mybackyard.com'

fork do
  # In this contrived example the parent process forks a child to take
  # care of sending data to the stats collector. Meanwhile the parent
  # process has continued on with its work of sending the actual payload.

  # The parent process doesn't want to be slowed down with this task, and
  # it doesn't matter if this would fail for some reason.
  StatsCollector.record message, recipient
end

# send message to recipient
```

+ For most other use cases involving fork(2) you’ll want some way to keep tabs on your child processes.

+ In Ruby, one technique for this is provided by Process.wait.
    + man 2 wait

+ Let’s rewrite our orphan-inducing example from the last chapter to perform with less surprises.

```ruby
fork do
  5.times do
    sleep 1
    puts "I am an orphan!"
  end
end

# Process.wait is a blocking call instructing the parent process to wait for one of its child processes to exit before continuing.
Process.wait
abort "Parent process died..."
```

+ If you have a parent that’s babysitting more than one child process and you’re using Process.wait, **you need to know which one exited**. For this, you can use the return value.

```ruby
# We create 3 child processes.
3.times do
  fork do
    # Each one sleeps for a random amount of number less than 5 seconds.
    sleep rand(5)
  end
end

3.times do
  # We wait for each child process to exit and print the pid that
  # gets returned.
  puts Process.wait
end
```

+ Process.wait has a cousin called Process.wait2
    + Process.wait  returns 1 value:  pid
    + Process.wait2 returns 2 values: pid, status

```ruby
# watch -d -n1 "ruby ruby/working-with-unix-processes/14.processes_can_wait/tmp.rb"

# We create 5 child processes.
5.times do
  fork do
    # Each generates a random number. If even they exit
    # with a 111 exit code, otherwise they use a 112 exit code.
    if rand(5).even?
      exit 111
    else
      exit 112
    end
  end
end

5.times do
  # We wait for each of the child processes to exit.
  pid, status = Process.wait2

  # If the child process exited with the 111 exit code
  # then we know they encountered an even number.
  if status.exitstatus == 111
    puts "#{pid} encountered an even number!"
  else
    puts "#{pid} encountered an odd number!"
  end
end
```

+ Process.wait has two more cousins called Process.waitpid and Process.waitpid2
    + Waiting for Specific Children
    + rather than waiting for any child to exit they only wait for a specific child to exit, specified by pid

```ruby
favourite = fork do
  exit 77
end

middle_child = fork do
  abort "I want to be waited on!"
end

pid, status = Process.waitpid2 favourite
puts status.exitstatus
```

+ Question: What if I haven’t gotten back around to Process.wait and another process exits? Let’s see:
    + The kernel queues up information about exited processes so that the parent always receives the information in the order that the children exited
    + So even if the parent is slow at processing each exited child it will always be able to get the information for each exited child when it’s ready for it
```ruby
# We create two child processes.
2.times do
  fork do
    # Both processes exit immediately.
    abort "Finished!"
  end
end

# The parent process waits for the first process, then sleeps for 5 seconds.
# In the meantime the second child process has exited and is no
# longer running.
puts Process.wait
sleep 5

# The parent process asks to wait once again, and amazingly enough, the second
# process' exit information has been queued up and is returned here.
puts Process.wait
```

+ Note that calling any variant of Process.wait when there are no child processes will raise Errno::ECHILD

+ It's always a good idea to keep track of how many child processes you have created so you don't encounter this exception.

+ The idea of looking in on your child processes is at the core of a common Unix programming pattern
    + babysitting processes
    + master/worker
    + preforking

+ Core of this pattern
    + you have one process that forks several child processes, for concurrency
    + and then spends its time looking after them: making sure they are still responsive, reacting if any of them exit, etc.




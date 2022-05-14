https://workingwithruby.com/wwup/zombies/

## Zombie Processes

+ fire and forget manner.

+ We need to revisit that example and ensure that we clean up that child process appropriately, lest it become a zombie!

## Good Things Come to Those Who wait(2)

+ The kernel will retain the status of exited child processes until the parent process requests that status using Process.wait.
    + If the parent never requests the status then the kernel can never reap that status information
    + So creating fire and forget child processes without collecting their status information is a poor use of kernel resources.

+ **If you’re not going to wait for a child process to exit using Process.wait (or the technique described in the next chapter) then you need to 'detach' that child process.**

```ruby
message = 'Good Morning'
recipient = 'tree@mybackyard.com'

pid = fork do
  # In this contrived example the parent process forks a child to take
  # care of sending data to the stats collector. Meanwhile the parent
  # process has continued on with its work of sending the actual payload.

  # The parent process doesn't want to be slowed down with this task, and
  # it doesn't matter if this would fail for some reason.
  StatsCollector.record message, recipient
end

# This line ensures that the process performing the stats collection
# won't become a zombie.
Process.detach(pid)
```

+ `Process.detach(pid)` simply spawns a new thread whose sole job is to wait for the child process specified by pid to exit. This ensures that the kernel doesn’t hang on to any status information we don’t need.

```ruby
# kernel/common/process.rb#L377-395
def self.detach(pid)
  raise ArgumentError, "Only positive pids may be detached" unless pid > 0

  # The evented system does not need a loop
  Thread.new { Process.wait pid }
end
```

+ What Do Zombies Look Like?
```ruby
# Create a child process that exits after 1 second.
pid = fork { sleep 1 }
# Print its pid.
puts pid

puts `ps aux #{pid}`

# Process.wait
```

+ So every child process that dies while its parent is still active will be a zombie, if only for a short time.

+ Once the parent process collects the status from the zombie then it effectively disappears, no longer consuming kernel resources.

+ It’s fairly uncommon to fork child processes in a fire and forget manner, never collecting their status. **If work needs to be offloaded in the background it’s much more common to do that with a dedicated background queueing system.**



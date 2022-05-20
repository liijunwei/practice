https://workingwithruby.com/wwup/zombies/

## Zombie Processes

+ fire and forget manner.

+ We need to revisit that example and ensure that we clean up that child process appropriately, lest it become a zombie!

## Good Things Come to Those Who wait(2)

+ The kernel will retain the status of exited child processes until the parent process requests that status using Process.wait.
    + If the parent never requests the status then the kernel can never reap that status information
    + So creating fire and forget child processes without collecting their status information is a poor use of kernel resources.

+ **If you’re not going to wait for a child process to exit using Process.wait (or the technique described in the next chapter) then you need to 'detach' that child process.**

如果父进程不相等子进程执行完毕, 那么父进程应该`detach`那个子进程

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

`Process.detach(pid)`做了一件事: 开出一个线程, 这个线程只做一件事, 那就是等这个被detach的子进程执行完毕

这样能避免内核资源的浪费(unclear)

```ruby
# kernel/common/process.rb#L377-395
def self.detach(pid)
  raise ArgumentError, "Only positive pids may be detached" unless pid > 0

  # The evented system does not need a loop
  Thread.new { Process.wait pid }
end
```

## What Do Zombies Look Like?

```ruby
# Create a child process that exits after 1 second.
pid = fork { sleep 1 }
# Print its pid.
puts pid

puts `ps aux #{pid}`

# Process.wait
```

## In The Real World

+ So every child process that dies while its parent is still active will be a zombie, if only for a short time.

一个父进程fork出了子进程, 子进程退出后, 但是父进程还没有推出并且还没有用wait收集子进程的信息时, 这个子进程就变成了僵尸进程

+ Once the parent process collects the status from the zombie then it effectively disappears, no longer consuming kernel resources.

当父进程把子进程的状态信息用`wait`方法收集了以后, 子进程才会消失

+ It’s fairly uncommon to fork child processes in a fire and forget manner, never collecting their status. **If work needs to be offloaded in the background it’s much more common to do that with a dedicated background queueing system.**







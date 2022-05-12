https://workingwithruby.com/wwup/orphans/

+ When starting a process via a terminal, we normally have only one process writing to STDOUT, taking keyboard input, or listening for that Ctrl-C telling it to exit.

+ It’s good to know about this stuff because it’s actually very easy to create orphaned processes:

```ruby
fork do
  5.times do
    sleep 1
    puts "I'm an orphan process #{Process.pid}!"
  end
end

abort "Parent process died..."
```

+ Question: What happens to a child process when its parent dies?
    + The short answer is, nothing.
    + The operating system doesn’t treat child processes any differently than any other processes.
    + When the parent process dies the child process continues on; the parent process does not take the child down with it.

+ concept: "Daemon processes"
    + long running processes that are intentionally orphaned and meant to stay running forever



+ concept: "communicating with processes that are not attached to a terminal session"
    + through unix signals


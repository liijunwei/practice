# [Appendix: How Resque Manages Processes](https://workingwithruby.com/wwup/resque/)

## The Architecture

> Resque is a Redis-backed library for creating background jobs, placing those jobs on multiple queues, and processing them later.

+ The component that we’re interested in is the Resque worker. Resque workers take care of the ‘processing them later’ part.

## Forking for Memory Management

+ Resque workers employ fork(2) for memory management purposes.

[lib/resque/worker.rb@perform_with_fork](https://github.com/resque/resque/blob/9e5324c65f6bd123819e63f2c365492f7516fd46/lib/resque/worker.rb#L907)

+ word: dissect 解剖

+ [procline is Resque’s internal way of updating the name of the current process.](https://github.com/resque/resque/blob/9e5324c65f6bd123819e63f2c365492f7516fd46/lib/resque/worker.rb#L864)


+ $0 is one of Ruby's global variables. [From here](https://ruby-doc.org/core-2.3.1/doc/globals_rdoc.html)
    + $0 -- Contains the name of the script being executed. May be assignable.

```bash
pry

[1] pry(main)> $0
=> "pry"
[2] pry(main)> $0 = "demo"
=> "demo"
[3] pry(main)>

ps aux|grep pry
ps aux|grep demo
```


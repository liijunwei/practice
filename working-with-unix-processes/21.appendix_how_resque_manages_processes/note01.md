# [Appendix: How Resque Manages Processes](https://workingwithruby.com/wwup/resque/)

+ Ruby job queue, [Resque](https://github.com/resque/resque)

## The Architecture

+ Resque is a Redis-backed library for creating background jobs, placing those jobs on multiple queues, and processing them later.

+ The component that we're interested in is the Resque worker. Resque workers take care of the 'processing them later' part.

## Forking for Memory Management

+ Resque workers employ `fork(2)` for memory management purposes.

+ [lib/resque/worker.rb@perform_with_fork](https://github.com/resque/resque/blob/9e5324c65f6bd123819e63f2c365492f7516fd46/lib/resque/worker.rb#L907)

+ word: dissect 解剖

+ [procline is Resque's internal way of updating the name of the current process.](https://github.com/resque/resque/blob/9e5324c65f6bd123819e63f2c365492f7516fd46/lib/resque/worker.rb#L864)


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

## Why Bother?

+ Resque uses `fork(2)` to ensure that the memory usage of its worker processes don't bloat.

Resque使用`fork(2)`来确保其工作进程的内存使用率不会膨胀

+ Let's review what happens when a Resque worker forks and how that affects the Ruby VM.

+ Once the child process is finished with the job it exits, which releases all of its memory back to the OS to clean up. Then the original process can resume, once again with only the application environment loaded.

+ So each time after a job is performed by Resque you end up back at a clean slate in terms of memory usage. This means that **memory usage may spike when jobs are being worked on**, but it should always come back to that nice baseline.

## Doesn’t the GC clean up for us?

+ When the Ruby VM boots up it is allocated a certain block of main memory by the kernel. When it uses up all that it has it needs to ask for another block of main memory from the kernel.
    + ruby vm 启动时会分配一定的内存, 当vm用光了内存后, 需要向内核申请更多内存

+ Due to numerous issues with Ruby's GC (naive approach, disk fragmentation) it is rare that the VM is able to release a block of memory back to the kernel. So the memory usage of a Ruby process is likely to grow over time, but not to shrink. Now Resque's approach begins to make sense!
    + 但是因为ruby的gc不好用, 没法使用gc把用后的内存完整释放, 所以ruby进程的内存使用率会不停的升高
    + 此时resque 的 模型就有用了
    + parent process 只负责分发任务("Resque uses `fork(2)` to ensure that the memory usage of its worker processes don't bloat. ")
    + worker process 执行任务("This is where memory usage can go awry出错. "), 执行完以后进程退出
    + 此时parent process 继续执行, 并且内存里还是只有很少的信息, 不存在使用了大量内存之后不释放的问题
        + "So each time after a job is performed by Resque you end up back at a clean slate in terms of memory usage. This means that memory usage may spike when jobs are being worked on, but it should always come back to that nice baseline."

## [Rescue Introduction](https://github.com/resque/resque#introduction)

Resque is heavily inspired by DelayedJob (which rocks) and comprises three parts(读源码时的思路):

+ A Ruby library for creating, querying, and processing jobs
+ A Rake task for starting a worker which processes jobs
+ A Sinatra app for monitoring queues, jobs, and workers.


+ [`lib/resque/worker.rb@work` is the main workhorse method. Called on a Worker instance, it begins the worker life cycle.](https://github.com/resque/resque/blob/9e5324c65f6bd123819e63f2c365492f7516fd46/lib/resque/worker.rb#L233)

### [Rescue demo is fun to play](https://github.com/resque/resque/tree/master/examples/demo)

```bash
git clone git://github.com/resque/resque.git

cd resque/examples/demo

gem install sinatra
gem install resque
gem install mono_logger
gem install mono_logger

# start server
rake start

# open web page
open http://localhost:9292/

# start worker
QUEUE=default rake resque:work
```


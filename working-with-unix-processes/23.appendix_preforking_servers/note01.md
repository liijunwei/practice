# [Appendix: Preforking Servers](https://workingwithruby.com/wwup/prefork/)

+ At the core of all these projects is the preforking model. There are a few things about preforking that make it special, here are 3:
    1. Efficient use of memory.
    2. Efficient load balancing.
    3. Efficient sysadminning.

## Efficient use of memory

+ Preforking uses memory more efficiently than does spawning multiple unrelated processes.

+ The truth is that it does take some time to fork a process (it’s not instantaneous瞬间) and that there is some memory overhead for each child process. These values are negligible可以忽略不计的 compared to the overhead of booting many Mongrels. Preforking wins.

+ **Keep in mind that the benefits of copy-on-write are forfeited if you're running MRI. To reap these benefits you need to be using REE.**
    + Q: what is REE?
    + A: [Ruby Enterprise Edition](https://rvm.io/interpreters/ree)
    + REE builds on top of MRI Rubies, versions 1.8.X and later, to deliver an enhanced interpreter with many performance and memory optimizations, including common patch-sets such as MBARI.

## Efficient load balancing

### The Very Basics of Sockets(complex topic)

+ Efficient load balancing has a lot to do with how sockets work.

高效的负载均衡和socket的工作方式有紧密的联系

+ Sockets are at the very core of networking.

+ Using a socket involves multiple steps:
    1. A socket is opened and binds to a unique port
    2. A connection is accepted on that socket using accept(2)
    3. Data can be read from this connection, written to the connection, and ultimately the connection is closed. The socket stays open, but the connection is closed.

使用socket有这么几个步骤:
1. 开启一个socket, 并把它和一个端口绑定
2. 在socket上建立连接
3. 在socket上通过连接读取数据/写入数据, 最后连接被关闭, 但是socket仍然是打开状态(???)

+ Typically this would happen in the same process.
+ A socket is opened, then the process waits for connections on that socket.
+ The connection is handled, closed, and the loop starts over again.

+ Preforking servers use a different workflow to let the kernel balance heavy load across the socket.

+ Thanks to the way fork(2) works, when the master process forks worker processes **each worker process gets a copy of the open socket**.

+ Kernel ensures that one, and only one, process can accept each individual connection.
+ Even under heavy load the kernel ensures that the load is balanced and that only one process handles each connection.

## Efficient sysadminning

+ preforking server     -> only need to control master process
+ non-preforking server -> control all instances

## [Basic Example of a Preforking Server](https://workingwithruby.com/wwup/prefork/#basic-example-of-a-preforking-server)

+ [code snippet](./basic_preforking_server_demo.rb)

+ You can consume it with something like nc(1) or telnet(1) to see it in action.

``` bash
nc localhost 8080
telnet localhost 8080
```

+ `Process.waitall` is simply a convenience method around Process.wait.
+ It runs a loop waiting for all child processes to exit and returns an array of process statuses.
+ Useful when you don't actually want to do anything with the process status info, it just waits for the children to exit.


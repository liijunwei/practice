# [Appendix: Preforking Servers](https://workingwithruby.com/wwup/prefork/)

+ At the core of all these projects is the preforking model. There are a few things about preforking that make it special, here are 3:
    1. Efficient use of memory.
    2. Efficient load balancing.
    3. Efficient sysadminning.

+ Preforking uses memory more efficiently than does spawning multiple unrelated processes.

+ The truth is that it does take some time to fork a process (it’s not instantaneous瞬间) and that there is some memory overhead for each child process. These values are negligible可以忽略不计的 compared to the overhead of booting many Mongrels. Preforking wins.

+ **Keep in mind that the benefits of copy-on-write are forfeited if you're running MRI. To reap these benefits you need to be using REE.**
    + Q: what is REE?
    + A: [Ruby Enterprise Edition](https://rvm.io/interpreters/ree)
    + REE builds on top of MRI Rubies, versions 1.8.X and later, to deliver an enhanced interpreter with many performance and memory optimizations, including common patch-sets such as MBARI.

## he Very Basics of Sockets(complex topic)

+ Sockets are at the very core of networking.

+ Using a socket involves multiple steps:
    1. A socket is opened and binds to a unique port
    2. A connection is accepted on that socket using accept(2)
    3. Data can be read from this connection, written to the connection, and ultimately the connection is closed. The socket stays open, but the connection is closed.

+ Typically this would happen in the same process.
+ A socket is opened, then the process waits for connections on that socket.
+ The connection is handled, closed, and the loop starts over again.

+ Preforking servers use a different workflow to let the kernel balance heavy load across the socket.

+ Thanks to the way fork(2) works, when the master process forks worker processes **each worker process gets a copy of the open socket**.

+ Kernel ensures that one, and only one, process can accept each individual connection.
+ Even under heavy load the kernel ensures that the load is balanced and that only one process handles each connection.

## [Basic Example of a Preforking Server](https://workingwithruby.com/wwup/prefork/#basic-example-of-a-preforking-server)

[code snippet](./basic_preforking_server_demo.rb)

You can consume it with something like nc(1) or telnet(1) to see it in action.

``` bash
nc localhost 8080
telnet localhost 8080
```


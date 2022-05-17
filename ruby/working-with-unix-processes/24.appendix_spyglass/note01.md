# [Appendix: Spyglass](https://workingwithruby.com/wwup/spyglass/)

+ Spyglass project was written specifically to showcase Unix programming concepts.

## [Spyglass' Architecture](https://workingwithruby.com/wwup/spyglass/#spyglass-architecture)

+ Spyglass is a web server.
+ It opens a socket to the outside world and handles web requests.
+ Spyglass parses HTTP, is Rack-compliant, and is awesome.

## booting spyglass

```bash
spyglass
spyglass -p other_port
spyglass -h # for help
```

## Before a Request Arrives

+ After it boots, control is passed to `Spyglass::Lookout`.
    + This class DOES NOT preload the Rack application and knows nothing about HTTP
    + it just waits for a connection.

## Connection is Made

+ When `Spyglass::Lookout` is notified that a connection has been made it forks a `Spyglass::Master` to actually handle the connection.

+ 当`Spyglass::Lookout`发现有连接进来时, 它会fork出一个`Spyglass::Master`去处理这个连接

+ `Spyglass::Lookout` uses `Process.wait` after forking the master process, so it remains idle until the master exits.

+ fork出master进程后, `Spyglass::Lookout`会wait阻塞这个进程, 直至master进程退出

+ `Spyglass::Master` is responsible for preloading the Rack application and forking/babysitting worker processes.

`Spyglass::Master`负责将Rack app加载进内存, 并且负责管理worker进程

+ The master process itself doesn’t know anything about HTTP parsing or request handling.

master进程不知道如何处理http请求

+ The real work is done in Spyglass::Worker

最终处理http请求的角色时是worker

 + It accepts connections using the method outlined in the chapter on preforking, leaning on the kernel for load balancing.

worker使用preforking接收连接, 依靠内核实现负载均衡

 + Once it has a connection it parses the HTTP request, calls the Rack app, and writes the response to the client.

处理完了连接请求后, 把响应返回给客户端

## Things Get Quiet

只要有新的请求进入应用, Spyglass就会不断重复前面建到的preforking过程, 然后用worker处理请求;

如果最后没有请求了, master进程和worker进程都会退出, 直至有新的请求进入应用

## Getting Started

[Now go forth and read the code!](http://github.com/d-bot/spyglass.git)

And may the `fork(2)` be with you!





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

+ After it boots, control is passed to `Spyglass::Lookout`.
    + This class DOES NOT preload the Rack application and knows nothing about HTTP
    + it just waits for a connection.


## Connection is Made

+ When Spyglass::Lookout is notified that a connection has been made it forks a Spyglass::Master to actually handle the connection.
+ Spyglass::Lookout uses Process.wait after forking the master process, so it remains idle until the master exits.

+ 当`Spyglass::Lookout`发现有连接进来时, 它会fork出一个`Spyglass::Master`去处理请求
+ fork后, `Spyglass::Lookout`会wait阻塞(???没看明白)




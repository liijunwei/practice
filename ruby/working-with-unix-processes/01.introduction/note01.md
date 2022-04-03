+ I knew how to use computers, but I had no idea how they worked under the hood.

+ late bloomer 大器晚成

+ As I worked on this high-traffic production system I was presented with increasingly complex problems.

工作中接触到流量很大的生产环境系统后, 遇到了越来越复杂的问题

+ As our traffic and resource demands increased we had to begin looking at our full stack to debug and fix outstanding issues.

随着流量增加和问题的复杂度提升, 我们必须开始深入理解整个系统, 去排查和修复问题

+ By just focusing on the application code we couldn’t get the full picture of how the app was functioning.

但是仅仅关注与业务(应用)的代码, 没法获取系统的全貌

+ We had many layers in front of the application: a firewall, load balancer, reverse proxy, and http cache.

在应用程序前还有很多层的东西支撑着, 例如 防火墙/负载均衡器/反向代理/http缓存等等

+ We had layers that worked alongside the application: job queue, database server, and stats collector.

和应用一起工作的还有 消息队列/数据库系统/指标收集系统等

+ This book will teach you all you need to know about Unix processes, and that is guaranteed to improve your understanding of any component at work in your application.

本书会教你unix进程相关的必要知识

并且保证学完以后一定能提升你对应用和应用周边服务的理解

+ Through debugging issues I was forced to dig deep into Ruby projects that made use of Unix programming concepts. Projects like Resque and Unicorn.

通过调试, 我被逼着深入了解了一些使用unix编程理论和技巧的ruby项目(gem), 例如 `Resque` 和 `Unicorn`

+ These two projects were my introduction to Unix programming in Ruby.

这两个项目带我进入了用ruby进行unix编程的大门

+ 


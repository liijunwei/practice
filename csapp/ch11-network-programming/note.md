## 11.3 因特网连接

+ 一个连接 是由它两端的套接字地址唯一确定的
    + (client-address:client-port, server-address:server-port)

+ "粗略一致和能用的代码" (p652) MIT Dave Clarke

## 11.4 套接字接口

+ 套接字接口(socket interface)是一组函数, 他们和unix IO函数结合起来, 用以创建网络应用

### 11.4.1 套接字地址结构

+ 从linux内核角度看, 一个套接字就是通信的一个端点
+ 从linux程序的角度来看, 套接字就是有相应描述符的打开文件(???)

### 11.4.2 socket函数

+ 客户端和服务器使用 socket 函数来创建一个套接字描述符(socket descriptor)

```bash
man 2 socket
man 2 connect
man 2 bind
man 2 listen
man 2 accept
```

+ 服务器调用 listen 函数告诉内核, 描述符是被服务器 而不是 客户端使用的

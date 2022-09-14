## 11.3 因特网连接

+ 一个连接 是由它两端的套接字地址唯一确定的
    + (client-address:client-port, server-address:server-port)

+ "粗略一致和能用的代码" (p652) MIT Dave Clarke

## 11.4 套接字接口

+ 套接字接口(socket interface)是一组函数, 他们和unix IO函数结合起来, 用以创建网络应用

### 11.4.1 套接字地址结构

+ 从linux内核角度看, 一个套接字就是通信的一个端点
+ 从linux程序的角度来看, 套接字就是有相应描述符的打开文件(???)


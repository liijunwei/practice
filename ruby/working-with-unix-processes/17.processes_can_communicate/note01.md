# [Processes Can Communicate](https://workingwithruby.com/wwup/ipc/)

+ This is part of a whole field of study called Inter-process communication(进程间通信)(IPC for short).
    + IPC implies communication between processes running on the **same** machine.

+ There are many different ways to do IPC, here covers two commonly useful methods
    + pipes
    + socket pairs

## Our First Pipe

+ A pipe is a uni-directional stream of data(单向的数据流)

```ruby
reader, writer = IO.pipe
```

+ The IO objects returned from IO.pipe can be thought of something like anonymous files.
    + You can basically treat them the same way you would a File. You can call #read, #write, #close, etc.
    + But this object won’t respond to #path and won’t have a location on the filesystem.

```ruby
reader, writer = IO.pipe
writer.write("Into the pipe I go...")
writer.close
puts reader.read
```

+ close the writer after writing to the pipe
    + because when the reader calls IO#read it will continue trying to read data until it sees an EOF
    + This tells the reader that no more data will be available for reading.
    + So long as the writer is still open the reader might see more data, so it waits.
    + By closing the writer before reading it puts an EOF on the pipe so the reader stops reading after it gets the initial data.

如果不关闭writer, reader认为还会有数据从管道传过来, 直到遇到EOF才会停止

## Pipes Are One-Way Only

+ The reader can only read
+ The writer can only write

```ruby
reader, writer = IO.pipe
reader.write("Trying to get the reader to write something") # will raise an IOError
```

管道是单向的, reader只能读, writer只能写

## Sharing Pipes

+ when a process fork a child, open resources are shared or copied to its child process.
+ Pipes are considered a resource, they get their own file descriptors and everything, so they are shared with child processes.

进程A fork出 进程B后, A的**资源**会和B进行共享
管道也是一种资源, 因此, 管道可以和子进程进行共享, 因此能用来进行父子进程间的通信

+ Here’s a simple example of using a pipe to communicate between a parent and child process.
```ruby
# (reader writer) x (parent child) = 4 endpoints
# we need to close two of them to(unused ends of pipe) so as not to interfere with EOF being sent
reader, writer = IO.pipe

fork do
  reader.close # child process use pipe's writer endpoint

  10.times do
    # heavy lifting
    writer.puts "Another one bites the dust"
  end
end

writer.close # parent process use pipe's reader endpoint
while message = reader.gets
  $stdout.puts message
end
```

父进程有两个管道的端点
子进程也有两个管道的端点(因为是fork出的, 被复制了, 但是指向的资源实际上是一样的)

其中, 因为父进程只需要读, 子进程只需要写(即子进程通过管道给父进程发消息), 所以把父进程的writer端点和子进程的reader端点关闭了

+ Since the ends of the pipe are IO objects we can call any IO methods on them, not just #read and #write.

+ Streams VS Messages
    + When I say stream I mean that when writing and reading data to a pipe **there’s no concept of beginning and end**.

+ Here’s a slightly more complex version of the pipe example where the child process actually waits for the parent to tell it what to work on, then it reports back to the parent once it’s finished the work:
```ruby
require 'socket'

child_socket, parent_socket = Socket.pair(:UNIX, :DGRAM, 0)
maxlen = 1000

fork do
  parent_socket.close

  4.times do
    instruction = child_socket.recv(maxlen)
    child_socket.send("#{instruction} accomplished!", 0)
  end
end
child_socket.close

2.times do
  parent_socket.send("Heavy lifting", 0)
end
2.times do
  parent_socket.send("Feather lifting", 0)
end

4.times do
  $stdout.puts parent_socket.recv(maxlen)
end
```

+ So whereas pipes provide uni-directional communication, a socket pair provides bi-directional communication. The parent socket can both read and write to the child socket, and vice versa.

+ Remote IPC?
    + ipc -> communication between processes running on the same machine
    + communicate via TCP sockets
    + communicate via RPC(remote procedure call)
    + ZeroMQ
    + ...

+ Both pipes and socket pairs are useful abstractions for communicating between processes.
+ They’re fast and easy.
+ **They’re often used as a communication channel instead of a more brute force approach such as a shared database or log file.**

+ As for which method to use: it depends on your needs. Keep in mind that pipes are uni-directional and socket pairs are bi-directional when weighing your decision.

+ System Calls
    + IO.pipe     maps to pipe(2)
    + Socket.pair maps to socketpair(2)
    + Socket.recv maps to recv(2)
    + Socket.send maps to send(2).




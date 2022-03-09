# [Processes Can Communicate](https://workingwithruby.com/wwup/ipc/)

+ This is part of a whole field of study called Inter-process communication(进程间通信)(IPC for short).

+ There are many different ways to do IPC, here covers two commonly useful methods
    + pipes
    + socket pairs

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


+ Pipes Are One-Way Only
    + The reader can only read
    + The writer can only write
```ruby
reader, writer = IO.pipe
reader.write("Trying to get the reader to write something") # will raise an IOError
```

+ when a process fork a child, open resources are shared or copied to its child process.
+ Pipes are considered a resource, they get their own file descriptors and everything, so they are shared with child processes.

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

+ Since the ends of the pipe are IO objects we can call any IO methods on them, not just #read and #write.

+ Streams VS Messages
    + When I say stream I mean that when writing and reading data to a pipe **there’s no concept of beginning and end**.






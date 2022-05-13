https://workingwithruby.com/wwup/cow/

+ CoW: Copy-on-Write semantics 写时复制

+ Physically copying all of that data(fork did this) can be considerable overhead, so modern Unix systems employ something called copy-on-write semantics (CoW) to combat this.

如果从父进程里物理复制数据可能会产生很大的开销

现代的unix系统采用了 "写时复制Copy on Write"的技术解决这个问题

+ CoW delays the actual copying of memory until it needs to be written.

写时复制技术会延迟内存的实际复制时机, 知道需要写入的时候才去复制

+ So a parent process and a child process will actually share the same physical data in memory until one of them needs to modify it, at which point the memory will be copied so that proper separation between the two processes can be preserved.

父进程和子进程实际上 在其一需要修改内存之前, 都是共享内存的

有一个需要修改内存的时候, 内存在会被复制和分隔开

```ruby
arr = [1,2,3]

fork do
  # At this point the child process has been initialized.
  # Using CoW this process doesn't need to copy the arr variable,
  # since it hasn't modified any shared values it can continue reading
  # from the same memory location as the parent process.
  p arr
end
```

```ruby
arr = [1,2,3]

fork do
  # At this point the child process has been initialized.
  # Because of CoW the arr variable hasn't been copied yet.
  arr << 4
  # The above line of code modifies the array, so a copy of
  # the array will need to be made for this process before
  # it can modify it. The array in the parent process remains
  # unchanged.
  puts "child: #{arr}"
end

puts "parent: #{arr}"
```

+ MRI's garbage collector uses a 'mark-and-sweep' algorithm.


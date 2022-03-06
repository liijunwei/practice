https://workingwithruby.com/wwup/cow/

+ CoW: Copy-on-Write semantics 写时复制

+ Physically copying all of that data(fork did this) can be considerable overhead, so modern Unix systems employ something called copy-on-write semantics (CoW) to combat this.

+ CoW delays the actual copying of memory until it needs to be written.

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


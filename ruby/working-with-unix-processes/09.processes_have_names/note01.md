https://workingwithruby.com/wwup/names/

+ There are two mechanisms that operate at the level of the process itself that can be used to communicate information
    + The process name
    + The exit codes

+ Every process on the system has a name
+ The neat thing about process names is that **they can be changed at runtime** and used as a method of communication.

```ruby
puts $PROGRAM_NAME

10.downto(1) do |num|
  $PROGRAM_NAME = "pry Process: #{num}"
  puts $PROGRAM_NAME
end
```


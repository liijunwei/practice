+ This plethora(过多/过量 [ˈpleθərə] ) of object creation is the result of the prior refactoring.

大量的对象创建是之前重构的结果

+ Over the course of the song, caching could reduce the number of new BottleNumber instances from 900 to 100.

如果使用实例变量, 可以把实例的创建数目由900降低到100

```ruby
def verse(number)
  bottle_number = BottleNumber.new(number)

  "#{bottle_number.quantity.capitalize} #{bottle_number.container} of beer on the wall, " +
  "#{bottle_number.quantity} #{bottle_number.container} of beer.\n" +
  "#{bottle_number.action}, " +
  "#{quantity(successor(number))} #{container(successor(number))} of beer on the wall.\n"
end
```

+ Now the first three phrases of the verse template send messages to a BottleNumber rather than to self. Only phrase four remains to be updated.



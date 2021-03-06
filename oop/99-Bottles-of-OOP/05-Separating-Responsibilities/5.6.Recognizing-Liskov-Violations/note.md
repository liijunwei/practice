```ruby
def verse(number)
  bottle_number = BottleNumber.new(number)

  "#{bottle_number.quantity.capitalize} #{bottle_number.container} of beer on the wall, " +
  "#{bottle_number.quantity} #{bottle_number.container} of beer.\n" +
  "#{bottle_number.action}, " +
  "#{quantity(successor(number))} #{container(successor(number))} of beer on the wall.\n"
end

def quantity(number)
  BottleNumber.new(number).quantity
end

def successor(number)
  BottleNumber.new(number).successor
end
```

+ Phrases 1 through 3 of the verse template refer to the same bottle number, and so can share the currently-cached BottleNumber instance.

+ Phrase 4, however, uses a different bottle number. Here’s a reminder of the code:

+ The plan is to change phrase 4 to send messages to instances of BottleNumber rather than to self.

+ This inconsistency is another violation of the generalized Liskov Substitution Principle.

里氏替换原则: 子类型(subtype)必须能够替换掉他们的基类型(base type)

+ A method named successor implicitly promises that the thing it returns will behave like the object to which you sent the message.

名为`successor`的方法暗含了它的返回值会有它的调用者一样的性质, 但是这个`BottleNumber#successor` 和 `Bottle#successor` 没有达到这种效果

+ It is often better to finish horizontal refactorings before undertaking vertical tangents.

> "Stay the course" is a phrase used in the context of a war or battle meaning to pursue a goal regardless of any obstacles or criticism.
>
> "坚持到底"是在战争或战斗中使用的短语, 意思是不顾任何障碍或批评, 追求目标

+ TODO fix the Liskov violation

+ You’ve already declared a temporary variable to hold bottle number 99. The current problem can be solved by declaring another variable to hold bottle number 98 and writing some shameless code.

临时方案: 使用一个新的BottleNumber实例, 存储`98`





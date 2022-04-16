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

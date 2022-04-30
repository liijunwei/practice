+ The final unresolved issue is that the lyrics instance method in BottleVerse directly references the concrete BottleNumber class.

```ruby
class BottleVerse
  # ...
  def lyrics
    bottle_number = BottleNumber.for(number)

    "#{bottle_number} of beer on the wall, ".capitalize +
    "#{bottle_number} of beer.\n" +
    "#{bottle_number.action}, " +
    "#{bottle_number.successor} of beer on the wall.\n"
  end
end
```

+ problems
    1. `bottle_number = BottleNumber.for(number)` 依赖于了 `BottleNumber`类 和 `for`方法, 如果它不需要知道这些最好
    2. `lyrics`方法里包含了一个空行; "Blank Line code smell" 提示我们, 这个方法里可能包含了多于一个职责, 违背了单一职责原则
    3. `lyrics`方法里只包含了一处对`number`的引用

+ Well-designed object-oriented applications consist of loosely-coupled objects that rely on polymorphism to vary behavior.

设计良好的面向对象程序由一系列松耦合的对象组成, 这些对象依赖于多态来产生各种各样的行为

+ Experienced object-oriented programmers know that applications most easily adapt to the unknown future if they:(TODO unclear)
    1. resist giving instance methods knowledge of concrete class names, and
    2. seek opportunities to move the object creation towards the edges of the application.

+ leeway [ˈliːweɪ] 自由活动的空间, 回旋余地




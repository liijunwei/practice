+ This horizontal refactoring to remove the case statement is almost complete.

使用前面讲到的的水平重构的方法, 几乎把case都合并了

+ This next section resolves the final difference, and in so doing illustrates the deep power of the Flocking Rules to unearth unanticipated abstractions.

接下来的这一节将会把最后两个case合并了(通过挖掘概念实现)

这么做是为了说明"Flocking Rules"挖掘意外抽象的深层力量

```ruby
  def verse(number)
    case number
    when 0
      # ...
      "99 #{container(number-1)} of beer on the wall.\n"
    else
      # ...
      "#{quantity(number-1)} #{container(number-1)} of beer on the wall.\n"
    end
  end
```

+ However, just because the tests pass doesn’t mean that the abstraction is correct.

测试通过不意味着做出了正确的抽象

+ Ask yourself these two questions:
    1. What is the responsibility of the quantity method?
    2. Isthereawaytomakethefourthphrasesmorealike,evenifnotyet identical?


+ There’s something subtle about the difference above, such that the underlying concept is not immediately obvious.

+ And this, unfortunately, is a constant of programming life.

+ If you had perfect understanding, you’d write perfect applications.
+ Mostly, however, you’re stumbling around, suffering from insufficient information, seeing problems through a glass, darkly.





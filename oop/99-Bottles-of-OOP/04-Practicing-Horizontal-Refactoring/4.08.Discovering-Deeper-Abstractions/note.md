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

如果你对问题有全面/完整的理解, 那么你能写出完美的应用

但是, 多数时候, 你都是在跌跌撞撞的/信息不足的情况下, 透过黑色的玻璃镜看问题

+ When you’re confused, don’t try to solve the entire problem straightaway.

当感到困惑的时候, 不要试图一下解决所有问题

+ The more confused you are, the more important it is to nibble.

越是感到困惑时, 越需要细嚼慢咽

+ You already know that it becomes easier to see how things are different if you make them more alike.

如果你能让代码里的case变得更加相似, 你就能更容易地看出他们之间的不同之处

+ Instead of trying to understand everything at once, simply search for a way to make line 5 above look more like line 8 (even if not identical), using existing code.

不要妄想一下子理解所有内容

一步一步思考

+ If you just realized that you can make these lines a little bit more alike by passing the 99 into quantity , you’ve got it.


```ruby
  def verse(number)
    case number
    when 0
      # ...
      "#{quantity(99)} #{container(number-1)} of beer on the wall.\n"
    else
      # ...
      "#{quantity(number-1)} #{container(number-1)} of beer on the wall.\n"
    end
  end
```

+ Having made these lines as similar as possible, it is now obvious that:

```ruby
"99"
```

must represent the same concept as:

```ruby
number-1
```

+ **As always, you must name this concept, create a method, and send the message in place of the difference.**

+ TODO 看明白 `String#succ` 方法是什么意思
```
Returns the successor to str.

The successor is calculated by incrementing characters starting from the rightmost alphanumeric (or the rightmost character if there are no alphanumerics) in the string.

Incrementing a digit always results in another digit, and incrementing a letter results in another letter of the same case.

Incrementing nonalphanumerics uses the underlying character set’s collating sequence.
```

```ruby
"a".succ # => "b"
9.succ   # => 10
```



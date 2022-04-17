+ The code now consists of a pleasing set of small objects with clear-cut responsibilities.

当前版本的代码已经由一系列小的对象组成了, 各个小模块职责清晰

+ However, there’s one persistent problem that can no longer be ignored: the successor methods violate the generalized Liskov Substitution Principle.

但是还有个问题需要处理: `successor`方法违反了里氏替换原则

+ You have every right to expect the successor of a bottle number to act like a bottle number

我们有理由认为 一个bottle_numer的实例的successor也具有bottle_number的行为

但是当前的`successor`方法返回了一个integer类型的数字, 显然它并不具有bottle_number的行为

+ Liskov violations are insidious, and over time cause increasing harm.

**违反 "氏替换原则" 是阴险的, 随着时间的推移, 违反这个原则所造成的伤害会越来越大**

+ If successor obeyed Liskov, you could substitute the hypothetical code on line 6 below for the code on line 5

```ruby
  def verse(number)
    bottle_number      = bottle_number_for(number)
    next_bottle_number = bottle_number_for(bottle_number.successor)
    # next_bottle_number = bottle_number.successor

    "#{bottle_number} of beer on the wall, ".capitalize +
    "#{bottle_number} of beer.\n" +
    "#{bottle_number.action}, " +
    "#{next_bottle_number} of beer on the wall.\n"
  end
```

如果遵循了里氏替换原则, 那么应该可以用`next_bottle_number = bottle_number.successor` 替换 `next_bottle_number = bottle_number_for(bottle_number.successor)`





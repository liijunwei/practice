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

在第四章的时候, 我们用"Flocking Rules"分析后提取出`successor`方法, 那时候它返回数字, 没有问题, 也没有违反里氏替换原则

在第五章的时候, 我们发现对原始数据依赖过多, 然后提取出了"BottleNumber"类, 然后把方法挪入了BottleNumber类里

但是这时候, successor方法的实现还是返回一个整型数字, 这时候违反了里氏替换原则

+ Alterations are needed in several places. Ultimately:
    1. The factory should be located such that it is reachable by the successor methods
    2. the successor methods should invoke the factory
    3. the verse method should expect successor to return a bottle number

+ Step 1 is to put the factory within successor's reach.

+ the "echo chamber" effect

"回音室"效应

+ The trick to moving forward using one-line changes is to temporarily alter the factory to tolerate both kinds of input.

接下来要重构`successor`方法, 让他调用工厂方法, 返回BottleNumber的实例

但是无论怎么修改, 都会使得测试用例失败

为了能够每次只做一点修改, 并且测试用例不崩溃, 使用一点技巧: 使得工厂方法同时支持传入数字和传入BottleNumber的实例

+ Temporary variables that are used just once can be removed with the Inline Temp refactoring

只使用一次的临时变量可以通过内联临时重构删除

+ Trustworthy objects are a joy to work with because they always behave as you expect.

值得信赖的实例行为是工作的乐趣, 因为它们总是按照你的期望行事

+ Having successfully fixed the problem with successor, it’s time to return to the main issue at hand.

开始处理新需求!




```ruby
class Bottles
  # ...
  def verse(number)
    verse_template.new(number).lyrics
  end
end
```

+ line 5 above violates `the law of Demeter`

+ From the classes-are-just-objects perspective, this line of code is very much a LoD violation.

+ Sending new to verse_template returns a object that conforms to a different API than the original verse_template receiver.

+ Put another way:
    + The API of the object that responds to new is different from the API of the object that responds to lyrics.
    + While a class and an instance of that class are certainly both objects, they are very different kinds of things.

+ The rule for injecting dependencies is that you should inject the thing you want to talk to.

注入依赖的规则是: 你只应该注入这个类想要直接交流的实例

+ The practical effect of this rule is to prohibit the use of injected objects in message chains that violate the Law of Demeter.

这样做的好处是 它禁止了在违反德米特法的消息链中使用注入对象

+ If new counts as a real message, and it does, then this rule also suggests that you should inject instances, not classes to which you are forced to send new and then something else—in this case, lyrics.

`new`方法也是一条消息

所以在`verse`里使用`verse_template.new(number).lyrics` 里面的 `new` 和 `lyrics` 就是对里氏替换原则的违背

这也指明你在注入依赖时, 应该注入一个实例, 而不是注入一个类

+ The rule gets broken here because in this case it is not possible to inject an instance of BottleVerse.

上面的例子里违背了迪米特法则, 因为这种情况下只能注入一个类, 没法注入一个实例(TODO unclear)

+ It’s long past time for these extracted classes to have their own tests.

+ Now just imagine how you might test the Bottles#verse method

+ Remember that testing is the first form of reuse, so this mental exercise will give you a hint about how easy it will be to reuse Bottles elsewhere in your application.

记得吗, 测试用例本身就是一类对原有代码的复用

所以这个心理练习将提示你在应用程序的其他地方复用这部分代码时 是简单还是困难

+ TODO "To make this exercise correspond to a real-life problem, pretend that lyrics is an expensive operation..." 没看懂(8.7.Identifying-What-The-Verse-Method-Wants p214)

+ Pain in testing is a sign of a rigid application and an indication that there’s something wrong with the design.

+ better desigin?
    + 想要这种消息直接要就好了...
    + 不必非要借助类的实例, 类本身就是Class的实例, 是这个思路吗
```ruby
# Listing 8.26: Collaborate Directly With Verse Template
class Bottles
  # ...
  def verse(number)
    verse_template.lyrics(number)
  end
end
```

+ If you want something from an object, just ask it. If it doesn’t know how to comply, teach it.

如果你想从一个对象身上得到什么, 直接去要就好了; 如果它不知道如何遵守(实现), 教它就好了

+ Don’t be trapped by what that other object currently does, but instead, loosen coupling by designing a conversation that embodies what the message sender wants.

不要被其他对象当前所做的事情所束缚

通过设计一个体现消息发送者想要什么的对话来降低耦合

(体会到了一点 实例方法 和 类方法 之间的一点差别)

```ruby
class BottleVerse
  # ...
  def self.lyrics(number)
    new(number).lyrics
  end
end
```

+ This class method does two things:
    1. it uses forwarding to eradicate the extra hop and resolve the Demeter violation
    2. it establishes a new API for players of the verse template role. Objects that play this role must respond to lyrics(number).

+ Inside of Bottles, the `verse_template` variable is now thought of as `holding a player of the verse template role` rather than the BottleVerse class.

+ The instance methods of BottleVerse are now private implementation details and have effectively disappeared from sight. Now that the BottleVerse.lyrics(number) class method exists, you might be tempted to move the behavior from those private instance methods into this new class method. This would certainly work, and it would reduce the amount of code. Why bother to create an instance of BottleVerse at all?

定义了类方法`lyrics`, 需求得到了实现, 实例方法`lyrics`可以变成私有的了

这时候你肯定会想, 为什么还需要这个实例方法呢?

虽然前面说过, 类也是Class类的实例, 但是 类 和 实例在面向对象的视角上看是不同

从实例的角度看, 行为和数据结合起来完成应用里的某个功能

Putting domain behavior on the class side rather than on the instance side places a bet that this domain concept will never involve data that varies. This bet makes sense only if the value of not typing n e and w today is greater than the future cost of converting all the class methods to instance methods should you find that data needs to vary.

将领域的问题放到类上, 而不是放到实例上, 就是在打赌: "这个域概念永远不会涉及变化的数据"

这种假设有时候会成立, 但是更多时候是不成立的;

当不成立时, 需要将实现在类上的方法转移到实例上, 代价很高

+ Since [class methods resist refactoring](https://codeclimate.com/blog/why-ruby-class-methods-resist-refactoring/), the cost of moving domain behavior from the class to the instance can be very high, and can far outweigh the paltry savings you get from avoiding typing new.

将领域的行为由类移到实例上的代价非常高, 常常会远超多些一个实例方法的代价

+ Because you cannot know the future, you cannot correctly guess when to follow which strategy. This suggests that the most economical overall plan is to always create instances of objects. Don’t waste a minute thinking about whether or not to do this. Part of your programming aesthetic is to reflexively put domain behavior on instances.

你无法预知未来, 最经济的的方法就是 创建一个类方法, 在类方法里创建实例, 再去调用实例方法

甚至不用花时间去想 是否需要这么做, 这么做就是了...

**编程美学的一部分是 将域行为本能地放在实例上**

+ Not only has the BottleVerse class has been extracted and injected, but the resulting Demeter violation has been fixed by adding a forwarding method.

+ Now that you understand the Law of Demeter, following it should become a part of your programming aesthetic. Be extremely biased towards fixing violations.

既然你了解了德米特法则, 遵循它应该成为你编程美学的一部分

一定要注意 纠正对迪米特法则有所违背的代码和设计

+ The overall cost of dealing with each transgression as it occurs is guaranteed to be less than the ultimate cost of repairing the few that spiral out of control after infecting your code for several years.

每次发现违背规则后立即就成的总成本, 一定会低于 积累了很多问题, 但是只在少数几次爆发后再去纠正的总成本

(功夫花在平时, 比堆积在一起去处理来的成本更低)

+ In general, don’t violate the Law of Demeter, and fix violations as you come upon them.

有些场景下违反迪米特法则是可接受的(数据库的例子, 没看懂...)

**但是一般来说, 一定要注意不要违反迪米特法则, 发现违背的地方是, 要及时处理**





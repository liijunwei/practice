# 8.6. Obeying the Law of Demeter

+ Lines that contain many dots might violate the Law of Demeter (LoD).
+ This section defines that law, determines where it applies, explores the consequences of ignoring it, and explains how to fix violations.

+ Listing 8.19: Verse Method Contains Many Dependencies
```ruby
class Bottles
  # ...
  def verse(number)
    verse_template.new(number).lyrics
  end
end
```

+ the verse method knows:
    + that verse_template responds to new
    + that new expects an argument
    + that the argument to new must be a number
    + that the object returned from new(number) responds to the message lyrics that lyrics returns the actual lyrics of interest

+ **This list enumerates many things that Bottles knows about but doesn’t control, which means they’re dependencies**
+ Dependencies are vulnerabilities

依赖 是 脆弱性的来源

+ Dependencies can’t be avoided but should certainly be minimized.

应用程序无法避免依赖, 但是 应该尽量减少依赖

+ Be alert for superfluous dependencies and remove them with extreme prejudice.

警惕多余的依赖, 并以极端偏见消除它们

+ [Law of Demeter](https://en.wikipedia.org/wiki/Law_of_Demeter) (TODO unclear)
    + In its general form, the LoD is a specific case of loose coupling.
    +
    + Each unit should have only limited knowledge about other units: only units "closely" related to the current unit.
    + Each unit should only talk to its friends; don't talk to strangers.
    + Only talk to your immediate friends.

迪米特法则, 是松耦合的一种特例

```ruby
class Foo
  def durability_of_preferred_toy_of_best_friends_pet
    best_friend.pet.preferred_toy.durability
  end
end
```

![Figure 8.3: Violating the Law of Demeter](./Xnip2022-04-27_09-18-55.png)

`Foo` 首先将 `pet`消息发送给了 `best_friend`, 因为Foo知道 `best_firedn` 的存在, 所以`Friend`是直接的协作者, 发送消息后, 返回了一个`Pet`的实例

目前为止没问题

接下来, `Foo`给`pet`实例发送了`preferred_toy`消息, 注意: `Pet` 不是`Foo`的直接协作者

`Pet`是`Friend`的直接协作者

它的意思是: `Foo`依赖了`Friend`的直接协作者, 依赖了自己协作者的协作者

接下来这个过程又重复了一次: `Foo`给`Toy`的实例发送了`durability`消息, 发送这条消息, 要求`Foo`了解`Toy`的api, 即






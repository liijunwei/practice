+ Switch your attention to the BottleNumber class. It’s full of conditionals, all of which have the same shape.(Code Smell)

+ vociferously 大声地(voice)

```ruby
# depend on the number argument
class Bottles
  def container(number)
    if number == 1
      "bottle"
    else
      "bottles"
    end
  end
end

# depend on the number message
class BottleNumber
  attr_reader :number

  def initialize(number)
    @number = number
  end

  def container
    if number == 1
      "bottle"
    else
      "bottles"
    end
  end
end

```

+ Extracting BottleNumber certainly removed the conditionals from Bottles, but they didn’t disappear: they just moved to the newly extracted class.

抽象出`BottleNumber`类虽然让`Bottles`类里的判断消失了, 代码好像更容易读了, 但是判断最终还是存在, 只是换了个地方

+ Fowler offers several curative refactoring recipes. The two main contenders are Replace Conditional with State/Strategy and Replace Conditional with Polymorphism.

Fowler 提出了几种重构的方法:
1. 用状态/策略 来替换条件判断
2. 用多态 来替换条件判断

+ dispersing 分散 疏散

+ The Replace Conditional with State/Strategy recipe removes conditionals by dispersing their branches into new, smaller objects, one of which is later selected and plugged back in at runtime. This recipe results in a code arrangement known as composition.

??? TODO

+ TODO: Replace Conditional with State/Strategy produces interesting results, and you are encouraged to experiment with that recipe on your own.

+ However, Replace Conditional with Polymorphism leads to a code arrangement that’s felicitous for the six-pack problem, and so will be followed in the next section.

+ The best way to figure out what will happen if you follow competing recipes is to try it.

+ Practice builds intuition. Do it enough, and you’ll seem magical too.






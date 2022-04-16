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


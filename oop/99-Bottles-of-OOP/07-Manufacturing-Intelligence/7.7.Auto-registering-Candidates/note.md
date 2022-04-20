+ Instead of requiring candidates to register themselves, you can make the factory figure out which classes ought to get registered and do it for them.

让新的类自己注册进列表里以外, 我们还可以让工厂替新的类做这件事

+ This requires that all candidates share some quality or characteristic that allows the factory to identify them.

这要求其他类都有某些性质或者特征, 工厂用它来做出识别

前面实现过这样的动态工厂
```ruby
class BottleNumber
  def self.for(number)
    begin
      const_get("BottleNumber#{number}")
    rescue NameError
      BottleNumber
    end.new(number)
  end
end
```

+ Common names aren’t the only thing that can enable the factory to find candidate classes.

一致的名字并不是工厂做出识别的唯一特征

+ Anything that makes candidates unique will work.

工厂能利用任何使得新类不同的特征做出识别

+ The key idea here is that there must be something about the candidate classes that allows them to be located by the factory.

+ stipulate [ˈstɪpjuleɪt] 规定

+ The *inherited hook* merits a bit of explanation.

得解释一下继承的hook

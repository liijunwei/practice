+ One final thought about BottleNumber6 before moving on: it may have occurred to you to meet the six-pack requirement by simply overriding to_s within BottleNumber6.

继续前进之前, 思考一下这种实现方式: 早BottleNumber6里覆写to_s方法

```ruby
class BottleNumber6 < BottleNumber
  def to_s
    "1 six-pack"
  end
end
```

这种实现方式也能满足需求, 但是仔细想想会发现它有问题

+ Consider the meaning of to_s versus that of quantity and container.

对比一下`to_s` 和 `quantity`+`container`方法的含义

+ The latter two methods reflect fundamental concepts in this domain. These concepts exist regardless of the way your application uses bottle numbers.

后者反映出了领域问题内的基本概念, 无论应用程序如何使用BottleNumber, 这些概念都存在

+ Extracting BottleNumber from Bottles decouples the idea of bottle number-ness from the "99 Bottles of Beer" song.

从Bottle类里提取出BottleNumber类, 将bottle_numer(瓶号)这个概念和瓶子解耦了

如果有需要, 在其他地方(非Bottle类里面)也可以使用BottleNumber的实例

+ Omitting quantity and container in favor of jamming "1 six-pack" directly into to_s corrupts BottleNumber6 with knowledge of the inner workings of the Bottles verse template.

如果忽略了`quantity`和`container`方法, 直接使用重写to_s方法实现了需求

就是使得使用BottleNumber6的地方必须知道内部细节, 破坏了抽象




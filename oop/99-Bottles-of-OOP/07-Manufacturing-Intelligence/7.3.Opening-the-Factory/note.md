+ BottleNumber0
+ BottleNumber1
+ BottleNumber6

这样的命名方式, 使得在工厂里面用元编程的方法动态获取类名变得可行

```ruby
  def self.for(number)
    begin
      const_get("BottleNumber#{number}")
    rescue NameError
      BottleNumber
    end.new(number)
  end
```

+ Even so, there are many things to dislike about this code. Here are a few common and thoroughly reasonable objections:

但是这个方法存在以下这些问题:

1. This version is harder to understand than the original.

元编程比switch case难懂

2. BottleNumber0, etc are no longer explicitly referenced in the source code.

具体的类不再直接被引用, 而是动态的被引用了, 搜索时可能不好查找

3. The code uses an exception for flow control.

代码里使用异常捕获做了流程控制

4. The factory ignores bottle number classes whose names do not follow the convention.

这个工厂会忽略所有没有按照命名规则创建的BottleNumber类

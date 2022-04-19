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


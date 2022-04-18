+ One final thought about BottleNumber6 before moving on: it may have occurred to you to meet the six-pack requirement by simply overriding to_s within BottleNumber6.

继续前进之前, 思考一下这种实现方式: 早BottleNumber6里覆写to_s方法

```ruby
class BottleNumber6 < BottleNumber
  def to_s
    "1 six-pack"
  end
end
```


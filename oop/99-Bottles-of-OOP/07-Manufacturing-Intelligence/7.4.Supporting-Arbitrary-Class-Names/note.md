+ The first step towards an open factory that both centralizes knowledge and supports arbitrary class names is to rearrange the code to increase the isolation of the names.

想要让工厂 既能将知识集中管理, 又让它对任意名字的扩展开放, 第一步需要做的是重新组织代码, 提升这些名字的隔离程度

+ You can do this by replacing the case statement with a key/value lookup, as follows:

可以把switch-case替换为键值对查找来实现

```ruby
class BottleNumber
  def self.for(number)
    Hash.new(BottleNumber).merge(
      0 => BottleNumber0,
      1 => BottleNumber1,
      6 => BottleNumber6)[number].new(number)
  end
end
```

+ The case statement factory is simple and allows arbitrary class names, but is closed. The meta- programmed factory is more complicated and requires a class naming convention, but is open. The key/value factory is similar to the case statement factory in that it allows arbitrary class names, but it’s a bit harder to read.

switch-case的版本很好懂, 但是它不对扩展开放

元编程的版本更加复杂, 但是它能够依赖命名的约定实现对扩展开放

键值对的版本和switch-case的版本类似, 并且它不依赖于命名的约定, 但是不容易读懂

+ This key/value version is slightly more complicated because the data has been separated from the algorithm. In this example, the "this⇒that" bits (the data) have been grouped together in one place (the hash) and the "if" bits (the algorithm) moved to another (the [] lookup logic).

其中键值对的版本更加复杂一些: 因为它把数据和算法隔离开了

数据部分放在了Hash(查找表)里面, 算法部分放在了`Hash#[]`里

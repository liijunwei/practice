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

+ The benefit of this separation is that you can now think of the driving data as an entity in itself, separate from the choosing algorithm.

数据和算法隔离的好处是, 你现在可以把数据视为一个整体, 把算法视为一个整体来思考

+ The algorithm lives in the code but you can store the data in an external file, or your database, and read it at initialization time. You might even create a nice user interface to update the database. You’ll have to update the database whenever a new class is added, but that’s a small price to pay for being able to change the behavior of your application without altering the actual code.

算法必须放在代码里, 但是数据确实可以隔离到代码外面的, 例如数据库, 在应用初始化的时候进行加载

甚至可以写一个UI界面, 来动态的更新数据库数据

这样, 在新加了类的时候, 只需要修改数据库(里的键值对)就可以了, 不需要修改工厂相关的代码

+ Pay some attention to the colors used in the syntax highlighting, When the colors change constantly it means that the code changes topics a lot.

+ Procedures are often characterized by many changes of color. Even if you are completely unfamiliar with this code, you can guess that the case statement factory is a procedure simply by looking at the alternating colors in the syntax highlighting.

过程式的代码通常能从代码着色上看出有很多主题变化

即使你完全不熟悉这段代码, 你也能通过颜色变化看出switch-case版本的工厂是过程式的

+ Code that is more object-oriented tends to group like things together, with fewer changes of topic. This results in more consistent colors as in the key/value factory.

OO的代码倾向于将相似的事情放在意思, 减少主题的转变

结果就是代码着色(语法高亮)看起来更加一致







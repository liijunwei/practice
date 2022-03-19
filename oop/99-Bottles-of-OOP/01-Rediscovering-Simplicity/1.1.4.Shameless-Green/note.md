+ TODO page 31

1. How difficult was it to write?

很容易

2. How hard is it to understand?

很容易

3. How expensive will it be to change?



+ Solving the 99 Bottles problem in any of these ways requires more effort than is necessary and results in more complexity than is needed.
+ These solutions cost too much; they do too many of the wrong things and too few of the right.

99 bottles problem的解法所花费的力气, 比想象中要大不少, 它的复杂度也没有看上去那么简单

之前看到的这些解法, 要么是太过"简洁", 要么太过追求"扩展性", 要么拆分过细

都没有找到一种把事情做对的做法

+ There’s nothing tricky here.

+ It feels embarrassingly easy, and is missing many qualities that you expect in good code. It duplicates strings and contains few named abstractions.

+ Most programmers have a powerful urge to do more, but sometimes it’s best to stop right here.

+ When you DRY out duplication or create a method to name a bit of code, you add levels of indirection that make it more abstract. In theory these abstractions make code easier to understand and change, but in practice they often achieve the opposite. One of the biggest challenges of design is knowing when to stop, and deciding well requires making judgments about code.

当通过抽象出方法去除冗余代码时, 你实际上向代码里注入了"间接性", 增加了代码的抽象程度

适度的抽象理应使得代码更易读, 更易改变, 但实际上却往往达到了相反的效果

设计的一个挑战就是: 知道什么时候该停止


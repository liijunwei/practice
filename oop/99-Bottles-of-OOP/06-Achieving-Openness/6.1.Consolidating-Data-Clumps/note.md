+ Despite much refactoring, the code is still not open to the six-pack requirement.

尽管前面做了很多重构工作, 但是现有代码还是没有能够对新需求的扩展"开放"

+ Once again, you must decide whether to continue forward with the existing code, or to retreat and strike out in a different direction.

此时需要再次确认, 我们想要继续重构, 直至它对新需求开放, 还是在现有代码基础上实现新需求

+ Consolidating Data Clumps

整合数据束(???)

+ Before focusing on that problem, however, there’s a simpler code smell that can be addressed.

处理问题前, 先把code smell消除掉

+ Data Clump is officially about data, and is defined as the situation in which several (three or more) data fields routinely occur together.

"Data Clump"是关于数据的官方术语, 意思是: 多个(3+)数据字段通常同时出现的情况

+ Having a clump of data usually means you are missing a concept.

代码出现了"data clump"通常意味着你忽略了一个概念(这个概念没被识别出来)

+ In the present case, it’s a slight stretch to call the quantity and container pairing above a Data Clump, If these two things always appear together, it’s a signal that this pairing represents a deeper concept, and that concept should be named.

在代码里能观察到这种模式: `quantity` 和 `container` 总是一起出现的

如果两个概念总是出现在一起, 它实际上是一个未被发现的更深层概念的信号, 并且我们应该把那个概念识别出来, 给他命名(提取方法)

+ Full-grown Data Clumps are usually removed by extracting a class, but in this small example it makes sense to simply create a new method.

+ It’s perfectly acceptable to override this default behavior, and many of your own classes would benefit from a custom to_s implementation.

BottleNumber类可以重写`to_s`方法

+ The verse method is getting simpler, but it still has more than one responsibility. This problem is reflected by the very structure of the code—the above method contains a blank line.

经过重构, `verse`方法已经被简化了很多了... 但是 它仍然有1个以上的职责 --- 看这个方法中间那个空行...🤣

+ Programmers add blank lines to signify changes of topic.

程序员通过添加空行来表示 代码做的事有了改变

+ The presence of multiple topics suggests the existence of multiple responsibilities, which makes code harder to understand when reading, and easier to harm when changing.



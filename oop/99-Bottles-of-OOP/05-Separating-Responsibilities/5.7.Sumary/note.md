+ This chapter continued the quest to make Bottles open to the six-pack requirement.

本章做的事是 继续重构Bottle类, 使得它 "对新需求所需的扩展开放"

+ It recognized that many methods in Bottles obsessed on number, and undertook the Extract Class refactoring to cure this obsession.

本章在改动前, 我们意识到很多方法(方法间的沟通)过于依赖基本数据类型(integer)number了, 这种依赖叫做: "Primitive Obsession", 是说应用程序太依赖于基本数据类型, 而不是依赖于抽象

+ The refactoring created a new class named BottleNumber.

然后从对基本数据类型的依赖里提取出了一个类`BottleNumber`

+ This chapter also explored the rewards of modeling abstractions, the trade-offs of caching, the advantages of immutability, and the benefits of deferring performance tuning.

+ 本章讨论了这些主题:
    + 建模是引入适当抽象的好处
    + 对数据进行缓存的利弊(权衡)
    + 数据不可变的优势
    + 延迟性能调整的好处, 不要过早(对性能进行)优化


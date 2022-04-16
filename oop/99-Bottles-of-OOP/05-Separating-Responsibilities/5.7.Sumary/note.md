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

+ The total ABC score score, for example, has gone up again.
    + `bash oop/99-Bottles-of-OOP/05-Separating-Responsibilities/test_all.sh`

ABC指标检测显示它的结果变得更糟糕了

+ Also, there are no unit tests for BottleNumber. It relies entirely on Bottle's tests.

并且新提取出的`BottleNumber`还没被测试覆盖到

+ The code still exudes many smells (duplication, conditionals, and temporary field, to name a
few).

当前代码还有很多code smell没有处理

+ And, finally, it commits a Liskov violation in the successor method.

并且还违反了里氏替换原则


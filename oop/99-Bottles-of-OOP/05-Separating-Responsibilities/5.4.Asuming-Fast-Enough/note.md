+ Immutability’s offsetting costs are twofold.
    + First, you must become reconciled to the idea, which for many programmers is no small thing.
    + Next, achieving immutability requires the creation of more (sometimes many more) new objects.

抵消不变性的成本是来自两方面:

1. 首先, 你必须接受(对象不可变)这个想法, 这对许多程序员来说并不是一件小事
2. 其次, 实现不变性需要创建更多的新对象

+ Getting habituated to a new way of thinking need happen only once, so this cost is not a permanent concern;

适应一种新的思维方式只需要发生一次, 它不是个大问题, 或者说它的阻碍并不大

+ The ongoing costs of immutability are therefore mostly in the creation of new objects, and that’s the topic of this section.

有问题的可能是接下来的使用"不变性"的持续成本---创建新的对象, 这是本节的主题

+ You may be familiar with Phil Karlton’s famous saying "There are only two hard things in Computer Science: cache invalidation and naming things."

你或许听说过Phil Karlton的名言: 计算机科学里只有两件难事, 1. 缓存失效 2. 命名

+ You’ve already read a great deal about naming things, and it’s finally time to discuss caching.

截止目前, 我们已经读了很多关于如何给变量/概念命名的思考, 接下来该讨论缓存相关的问题了

+ A cache, in computer science, is a local copy of something stored elsewhere.

计算机科学里, 缓存指的是 对存储在某处的变量的 局部副本(???)


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

+ cache are expected to:
    + make applications faster
    + lower costs

+ sometimes they are true, but not always

+ When you send a message and save the result into a variable, you’ve created a simple cache.
+ If the value in your variable becomes obsolete, you must invalidate this cache, either by discarding it, or by resending the message and saving the new result.

+ Caching is easy. However, figuring out that a cache needs to be updated can be hard.

为变量添加缓存是容易的, 但是决定缓存何时该被更新则是困难的

+ Notice that the costs of caching and mutation are interrelated.

请注意, "缓存变量值"和"改变变量值"的成本是相互关联的

+ Outdated caches can be a source of opaque, expensive, and frustrating bugs.

失效(过时)的缓存可能是复杂诡异的bug的来源

+ Given this, the best programming strategy is to write the simplest code possible and measure its performance once you’re done.

鉴于前面对缓存(可变性)和不可变性的考虑, 最好的编程策略是: 先写出最简单的代码, 然后度量它的性能(不要提前做出过于复杂, 以为会提升性能的设计)

+ Increasing speed may require caching, but many problems can be fixed by substituting more efficient code in specific, narrow places.

提升系统的性能不一定要通过添加缓存的方式实现; 或许优化系统里最慢的部分即可

+ Once you understand precisely what’s wrong, it may be possible to fix it without caching at all.

如果能准确定位到代码里慢的位置, 就很有可能通过优化代码的方式提升系统高性能(不需要添加缓存)

+ Your goal is to optimize for ease of understanding while maintaining performance that’s fast enough.

我们的目标应该定为 提升代码的易读性的同时保证系统性能

+ The first solution to any problem should avoid caching, use immutable objects, and treat object creation as free.

拿到问题后, 第一反应应该是避免使用缓存, 应该使用不可变的对象, 并且忽略创建对象的代价

+ This results in speedy development of simple code, which leaves plenty of time to identify and correct the real performance problems.

如果一这种思路思考, 结果是我们能很快开发出简单版本的代码, 然后留出很多时间鉴别性能瓶颈,然后针对性能瓶颈进行优化

+ Now that this somewhat theoretical discussion is complete, it’s time return to the Bottles class, and apply ideas to actual code.



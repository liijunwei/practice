+ This chapter pulled the lyrics of the "99 Bottles" song out of Bottles and put them into a new BottleVerse class.

本章将`lyrics`方法从`Bottles`类里提取出来, 组成了单独的类`BottleVerse`

It then injected an instance of BottleVerse back into Bottles.

然后把`BottleVerse`类的实例当做依赖注入到了`Bottles`类里, 使之和原来的功能等价

+ Extracting BottleVerse reduced the Bottles's responsibilities, making it easier to understand and maintain.

提取出`BottleVerse`类减少了`Bottles`类的职责, 使得`Bottles`更容易理解和维护

+ Injecting BottleVerse into Bottles loosened the coupling between Bottles and the outside world.

将`BottleVerse`注入到`Bottles`减少了`Bottles`和外界的耦合关系

+ Bottles now thinks of itself as being injected with players of the verse template role and will happily collaborate with any newly arriving object as long as that object responds to lyrics(number).

`Bottles`现在可以认为它是和一系列被注入的依赖(特定的玩家/角色)来交互, 不和某个实体交互

+ The impetus behind extracting BottleVerse was a new requirement to produce songs with different lyrics, that is, to vary the verse.

诗歌背后的推动力是一个新的需求, 即创作出不同歌词的歌曲

+ The refactorings in this chapter satisfied that requirement by following a fundamental strategy of object-oriented design: extracting the BottleVerse class isolated the behavior that the new requirement needed to vary.

本章中的重构, 通过遵循面向对象设计中的基本策略来满足该需求: 提取出`BottleVerse`类来隔离新需求所需要改变的行为

+ While continuing to lean on code smells and refactoring recipes, this chapter introduced the idea of a programming aesthetic.

在继续依靠"Code smell"和重构方法的同时, 本章介绍了编程美学的概念

+ A programming aesthetic is the set of internal heuristics that guide your behavior in times of uncertainty.

编程美学是一系列的"内部启发法", 能够在不确定该如何选择的时候知道我们的行为

+ This chapter suggested five precepts that belong in everyone’s object-oriented programming aesthetic:
    1. Put domain behavior on instances.
    2. Be averse to allowing instance methods to know the names of constants.
    3. Seek to depend on injected abstractions rather than hard-coded concretions.
    4. Push object creation to the edges, expecting objects to be created in one place and used in another.
    5. Avoid Demeter violations, using the temptation to create them as a spur to search for deeper abstractions.

1. 将领域的行为放到实例上(领域行为少用类方法实现)
2. 不要让实例方法知道常量的名称
3. 让类依赖于被注入的抽象(实例), 不要依赖于硬编码的具体实现
4. 将对象创建推到边缘, 要习惯于在一个位置创建对象 并在另一个位置使用(TODO unclear)
5. 避免违反迪米特法则, 通过挖掘引入违反这个法则的原因, 深入认识程序, 找到更深层的抽象 (*****)

+ The practical effect of following these precepts is to loosen the coupling between objects.

遵循这些规则的效果是 能够减少对象之间的耦合

+ Any application that survives will change.

应用程序必然会发生改变, 这是确定的

变化会发生在哪里, 这是不确定的

变化的确定性和不确定性提示我们, 最好的编程策略是 从最初创建的那一刻起, 努力较少所有代码之间的耦合

+ Instead of writing code that speculatively imagines a later need for one specific feature, they tell you to loosen the coupling of all code so that you can easily adapt to whatever future arrives.

不要编写哪些 "预测未来可能需要某些改变/某些功能"的代码, 而要从一开始就注意减少模块之间的耦合, 这样就能轻松地应对(适应)未来的变化







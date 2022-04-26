+ impetus [ˈɪmpɪtəs]  动力/势头

+ Recall that the impetus behind extracting BottleVerse was the need to produce songs with other lyrics.

回想我们从Bottles类里面抽取出BottleVerse类的动机是 满足为不同诗句产生不同歌词的需求

+ Despite having completed the extraction, you can’t yet fulfill this vary-the-verse requirement.

尽管目前已经实现了提取出新的类, 但是现在还是满足不了新需求

+ Why not? Because Bottles is currently stuck to BottleVerse. You have extracted the class, but not yet inverted the dependency.

为什么不行呢? 因为Bottles类现在被固定在了BottleVerse上(Hardcoded)

我们只抽取了类, 但是没有反转依赖关系

+ The Dependency Inversion Principle (DIP) contributes the 'D' in the SOLID acronym and can be defined as "depend on abstractions, not concretions."

依赖反转原则是说 我们的代码应该依赖于抽象, 而不应该依赖于具体的实现

## 8.5.1. Injecting Dependencies

+ It’s time to loosen the coupling between Bottles and BottleVerse.

+ Bottles depends on(has knowledge about) two different BottleVerse-related things:
    1. a concretion, that is, the name of the BottleVerse class, and
    2. an abstraction, namely, the idea that there’s an object that can provide a verse.

Bottles依赖于BottleVerse相关类的两部分

1. 类的具体的实现
2. 抽象的某个类的实例 能够产生诗句

+ Knowing the abstraction is required.

其中对抽象知识部分的依赖是不可避免的

+ Knowing about the concretion, on the other hand, is completely avoidable.


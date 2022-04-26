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

其中对抽象知识部分的依赖是必要且不可避免的

+ Knowing about the concretion, on the other hand, is completely avoidable.

但是对具体实现细节的依赖确实可以避免的

+ Bottles doesn’t have to know the concrete class name BottleVerse, this name could easily be passed into Bottles from the outside.

Bottles类不需要知道BottleVerse的具体实现, 它可以很容易的通关其他方式注入到Bottle类中

+ Doing so not only reduces the number of dependencies inside of Bottles, but it also opens Bottles to an entire universe of existing and potential lyrics providers.

这么做不仅可以减少Bottles内部的依赖, 并且使得Bottles变得对扩展开放

+ You can think of classes that provide lyrics as playing a common role.

可以将产生诗句的类视为充当某种角色

+ Roles need names, and this role could reasonably be named verse template.

这种角色我们给它起名为 "verse_template"

+ Question: when injecting collaborators, should you inject classes or instances of those classes?

## 8.5.2. Isolating Variants

+ Bottles now thinks of itself as interacting with a player of the verse template role rather than a kind of the BottleVerse type.

+ Bottles can produce as many different songs as you have different verse templates, without itself changing.

+ Here’s the process used to create the verse template role:
    1. Identify the code you want to vary.
    2. Name the underlying concept.
    3. Extract the identified code into its own class.
    4. Inject this new role-playing object back into the object from which it was extracted.
    5. Forward messages from the original class to the injected object.

意识到 verse_template 



+ To mutate is to change. State is "the particular condition of something at a specific time."

mutate即改变

状态(state)是"某事物在特定时间的特定状态"

+ The real world is pervaded(弥漫) by this idea--—what exists, will change.

真实世界中弥漫着这样一种观念：凡存在的东西都会改变

+ Human agreement about the necessity and rightness of change is reflected in the choice of the word variable for use within computer programming languages.

人们对改变的"必要性"和"正确性"的共识反映在计算机编程语言中变量这个词的选择上(???)

+ What purpose has a variable other than to vary? Most object-oriented programmers write code that both expects and relies upon object mutation. Objects are constructed, used, mutated, and then used again.

变量存在的目的除了"改变"以外, 还有什么呢?

多数使用面向对象思想编程的程序员写出的代码, 都依赖与实例的变化

实例被创建/使用/改变/重用

+ Regardless of how intuitive and natural it may seem, mutation is not an absolute requirement.

无论面向对象的思想看起来有多么的直观和自然, "可变性"都不是一个绝对的需求

+ It is perfectly possible (as programmers of functional languages will happily inform you) to construct applications from immutable objects, meaning objects that do not change.

使用不可变的对象也能构造出很好的应用(函数式编程)

+ Instead of refilling your existing cup, you discard it in favor of a new one that looks identical but is full of coffee.

如果你想要一杯咖啡, 你不需要装满现有的空杯子，而是把空杯子扔掉，换一个看起来一模一样但装满咖啡的新杯子

+ Rather than changing yourself to be more fit, you swap yourself for the new, fitter, you.

你不需要改变自己去适应, 你只需要把自己换掉, 换成一个新的"你", 一个更合适的你

+ As the Himalayas rise, you replace your existing copy with a brand new mountain range that’s a tiny bit taller.

如果喜马拉雅山脉升高了, 你会直接用一个稍高一点的全新"喜马拉雅山脉"替代现有的"喜马拉雅山脉"

+ Ponder the benefits of working with objects that do not change. What virtue might immutability provide, and what trouble might it avoid?

思考一下 使用不可变的对象有什么好处?

他能提供什么好处? 能避免什么麻烦?

+ One of the best things about immutable objects is that they are easy to understand and reason about.

不可变对象的最大好处之一是: 他们很容易理解和推理

+ Because they are easy to reason about, immutable objects are also easy to test.

由于不可变对象很容易推理, 对他们进行测试会变得容易

+ Another key virtue of immutable objects is that they are thread safe.

另一个关键的好处是: 不可变对象是线程安全的

多线程导致的问题极难排查和复现, 使用不可变对象, 能够避免这种情况

+ Having read this section, look back at the new BottleNumber class in "Listing 5.21: Forward Messages to Smarter Number". The question of mutability applies directly to this new class.

基于上面对可变对象和不可变对象的理解, 回头看看 `BottleNumber` 的实现

+ Imagine that you’re holding onto an instance of BottleNumber whose number variable contains the value 99.

想象你有一个 `BottleNumber` 的对象, `#<BottleNumber:0x00007fa224023140 number=99>`

+ The verse progresses such that it now needs bottle number 98. Is it better to mutate the value of number in the current instance of BottleNumber, or should that object be discarded in favor of BottleNumber.new(98)?

但是在打印诗句的时候, 每一句既用到了99, 也用到了98

这时候两种选择: 1. 修改实例变量number的值, 以适应99和98; 2. 创建一个新的对象, 专门表示98

+ **If you lean towards mutating the existing BottleNumber rather than making another, it’s possible that you are biased against creating new objects.**

如果你倾向于改变现有对象的状态而不是创建一个新对象的话, 那么很可能你是对创建新对象有偏见

+ **This bias is often unexamined, and has its roots in the assumption that if you routinely create many new objects, your application will be too slow.**

这种偏见通常是没有经过仔细检查和反思的, 它的根源在于一种假设: 加入你经常创建很多新的对象, 你的程序会变慢


1. How difficult was it to write?

难

2. How hard is it to understand?

难

3. How expensive will it be to change?

单独修改都容易, 但修改一个方法会影响到很多其他的方法

+ This solution is characterized by having many small methods.
+ This is normally a good thing, but somehow in this case it’s gone badly wrong.

这种解法使用了很多小的方法, 通常这种做法是有益的, 但这里确实错的


+ It’s obvious that the author of this code was committed to doing the right thing, and that they carefully followed the Red, Green, Refactor style of writing code.
+ The various strings that make up the song are never repeated —it looks as though these strings were refactored into separate methods at the first sign of duplication.

+ The code is DRY, and DRYing out code should save you money.
+ DRY promises that if you put a chunk of code into a method and then invoke that method instead of duplicating the code, you will save money later if the behavior of that chunk changes.

这些代码是dry的, 没有冗余, 减少代码的冗余能为你省钱

+ When you invoke a method instead of implementing behavior, you add a level of indirection.
+ This indirection makes the details of what’s happening harder to understand, but DRY promises that in return, your code will be easier to change.

当用调用一个方法(library)替代自己实现时, 实际上增加了一层间接性

这个间接性会让代码变得更难于理解

但是DRY带来的好处是 它使得代码更易于修改

+ The Don’t Repeat Yourself principle, like all principles of object-oriented design, is completely true.
+ However, despite that fact that the code above is DRY, there are many ways in which it’s expensive to change.

"不要重复" 原则和其他面向对象设计一样, 是[正确]的.
但是尽管这种实现方式是DRY的, 但是它仍然很难于修改, 因为它做得太过了

当需要把 "beer" 换为其他的实体时, 虽然只需要把 `beer` 方法的内容替换了就可以(因为是DRY的)
但是, 如果只替换方法的内容, 会发现方法名会变得非常奇怪
因为这些方法和`beer`有很深的绑定关系

应该做的是把beer的概念提取出来, 抽象一层

`the word "beverage" is one level of abstraction higher than "beer." `
例如: 饮料就是beer的抽象

+ The fault here, however, lies not with the DRY principle, but with the names of the methods.

这里出现的问题不再有DRY, 而在于方法的命名

+ Unfortunately, when you name a method after its current implementation, you can never change that internal implementation without ruining the method name.

如果方法的命名依赖于它的实现, 那么就没法只改变实现 并且不修改方法名(就能清晰表达含义)

+ You should name methods not after what they do, but after what they mean, what they represent in the context of your domain.

方法的命名不应该以他们做了什么命名, 而应该以他们的含义命名, 以他们代表的上下文来命名

+ Naming the method at this slightly higher level of abstraction isolates the code from changes in the implementation details.

为方法命名时, 稍微把抽象层次提高一点, 可以把代码的改动和他的具体实现隔离开

+ However, in Concretely Abstract, this force is overcome by the high cost of dealing with methods that are named at the **wrong level of abstraction**. These method names raise the cost of change.

一定要注意抽象的层次, 抽象的度

+ Therefore, one lesson to be gleaned from this solution is that you should name methods after the concept they represent rather than how they currently behave.

方法命名时, 应该依据它所代表的概念, 而不是它的行为

+ Many methods in this code represent the wrong abstractions.

方法定义的抽象粒度不合适, 也会让代码改变变得困难










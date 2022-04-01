## 3.7.1. Focusing on Difference

+ While it’s true that there are problems for which the solution is obvious, those of any interesting size aren’t tractable to instant understanding. They’re too big, or have too many parts.

虽然有的问题看起来解决方案显而易见

但是有趣的问题, 他们的复杂度使得他们不那么容易理解, 不容易解决

这些问题太大, 太复杂, 他们有太多的组成部分

+ When examining complicated problems, the eye is first drawn towards sameness.

当研究复杂的问题时, 眼睛首先会被吸引到相同的地方

+ However, despite the fact that sameness is easier to identify, difference is more useful because it has more meaning.

然而, 尽管相同性更容易识别, 但差异性更有用, 因为它具有更多意义

+ The "Gang of Four"
    + Erich Gamma
    + Richard Helm
    + Ralph Johnson
    + John Vlissides
    +
    + their joint authorship of [Design Patterns: Elements of Reusable Object-Oriented Software](https://www.amazon.com/Design-Patterns-Elements-Reusable-Object-Oriented/dp/0201633612)

+ This influential book describes **twenty-three patterns or solutions to common OO programming problems** and it explains this process thusly:

> The focus here is encapsulating the concept that varies, a theme of many design patterns.

+ Difference holds the key to understanding.

差异是理解代码的关键

+ If two concrete examples represent the same abstraction and they contain a difference, that difference must represent a smaller abstraction within the larger one.

如果两个实体代表了相同的抽象, 并且他们包含了差异, 那么这些差异一定代表了更高层次的抽象

+ If you can name the difference, you’ve identified that smaller abstraction.

如果你能识别出这些差异, 你就能识别出这些(小的)抽象

+ The good news is that a systematic application of the rules of refactoring converts difference to sameness, decomposing a problem into its constituent parts.

好消息是 系统地应用重构的法则 能将差异转换为相同, 能将问题分解为它的组成部分

+ The even better news is that this happens automatically.

更好的是, 这个过程是自动发生的(??)

+ You don’t have to identify the underlying abstractions in advance of refactoring. If you merely write the code dictated by the rules, the abstractions will follow.

你不必在开始重构之前就能准确分辨出底层的抽象

如果你按照这些法则去写代码, 这些抽象会自动浮现出来

+ The habit of believing that you understand the abstraction, and of jumping to an invented solution, is deeply ingrained. Programmers study a problem, decide on a solution, and then implement it. Solutions are crafted by intention.(????)

"相信自己理解了问题的抽象并直接实现了解决方案" 这种习惯已经根深蒂固了

程序员了解一个问题, 确定方案, 然后实现方案, 解决方案是由目的决定的(怎么翻译呢?)

+ If this describes your entire past experience, you may find the following code surprising. It takes many small, iterative steps, and results in a solution that is discovered by refactoring.

如果上面的话描述了你过去的经历, 你可能会觉得之后的代码会很让你吃惊

写这些代码的时候, 每次只做一点改动, 但是最终的解决方案是通过不断做小的重构得到的

> **First identify the things that are most alike**

选两个最相似的分支, 然后集中注意力把他们变成一样的

`case 1` 和 `case 2`(page 102) 唯一真正的区别是`bottle` 和 `bottles`

## 3.7.2. Simplifying Hard Problems

> **Next task is to make them identical.**

+ It’s important to focus on this specific goal without succumbing to the temptations of tangents.

当前, 最重要是专注于这个特定目标; 不要分心

+ Think of the process of turning these two lines into one as being on a horizontal path.

将这两条线合二为一的过程 可以看作是在水平路径上

在这条路径上的时候, 要抵御向垂直方向分析的诱惑(**目标以外皆为诱惑**)

+ If you begin making changes to other parts of the code before you completely combine the 2 and else cases, you step off a well-trod path into a woods so dark and sinister that you might never return.

如果你在完成将两个case合二为一之前, 开始修改其他部分, 那么你就是偏离了一条成熟的道路, 然后误入歧途

+ While it can be useful to interleave horizontal and vertical work, it’s best to finish the current journey when the terminus of the horizontal path is in sight.

在水平路径上努力是最有效的

```bash
cd $HOME/OuterGitRepo/1st_99bottles_ruby && git checkout chapter-3
```

+ don’t discount the value of solving easy problems.

+ After making the above change (and running the tests between each, of course), the remaining difference is "bottle/bottles" on the last line:

```ruby
when 2 # ...
"#{number-1} bottle of beer on the wall.\n" else
# ...
"#{number-1} bottles of beer on the wall.\n" end
```

+ This is the first interesting difference. Now you must decide what this difference means.

## 3.7.3. Naming Concepts

+ The goal of the current refactoring is to find a way to express that more abstract verse.

当前重构的目标是在找到一个能表达上面的差异的一种抽象

+ To make these two lines the same, you must name this concept, create a method named after the concept, and replace the two differences with a common message send.

为了让上面的两行变为一行, 需要给差异命名, 然后把它抽象为一个方法, 用这个方法替换那个差异

需要从 bottles 和 bottle 中提取出一个概念

+ "bottle" is not underlying the concept. If you call the method "bottle" you are naming it after its current implementation, and you’ve already seen how that can go badly wrong.

前面的例子已经说明过, "bottle"不是一个潜在的概念

如果你把抽象出的方法命名为"bottle", 那么一定是错的, 因为bottle是以它的实现命名的, 并没有做出抽象

+ Within the context of the song, "bottle/bottles" does not represent pluralization.

+ There are two pieces of information that can help in the struggle for a name. One is a general rule and the other is the new requirement.

有两条信息能帮助你做出命名的选择:

+ First, the new requirement. Recall that the impetus for this refactoring was the need to say "six-pack" instead of "bottle/bottles" when there are 6 bottles. The string "six-pack" is one more concrete example of the underlying abstraction. This suggests that if you name the method "bottle," you will regret this decision in short order.

1. 新的需求

新需求是在一定条件下把bottle替换为"six-pacs", 不是替换为"bottle/bottles"

"six-pack"是潜在抽象的 具体的例子

它暗示了, 如果你用"bottle"命名, 就错了

+ The general rule is that the name of a thing should be one level of abstraction higher than the thing itself. The strings "bottle/bottles/six-pack" are instances of some category, and the task is to name that category and to do so using language of the domain.

2. 通用的法则

通用的法则是: 命名的时候, 名字应该做出高于某个对象**一层**的抽象

"bottle/bottles/six-pack"都是某类事物的实例

在此基础上上升一层, 会是什么?

+ One way to identify the category is to imagine the concrete examples as rows and columns in a spreadsheet.

辅助方法: 列一个下面这样的表格

Number   | xxx?
---------|----------
  1      | bottle
  6      | six-pack
  n      | bottles

这里的xxx是什么比较合适呢?

+ It might seem as if "Unit" would be a good header.

很容易会想到一个命名: "unit"

+ Although it’s true that every example is some kind of unit, there are two problems with this name.

虽然Unit可以, 但是它并不合适, 它有这两个问题

+ First, it’s too abstract. Unit is not one level of abstraction higher than the examples, it’s many. There are plenty of good naming alternatives on the continuum between "bottle" and "unit."

1. 它太抽象了, 不是"one level of abstraction higher than the thing itself", 比它合适的命名还有很多

+ Next, unit is not in the language of the domain. The name you choose will be the name you use in conversations with your customers.


2. "unit"不是这个领域的语言; 选择的名字, 是要和客户进行沟通的

+ **Naming things after domain concepts improves communication between you and the folks who pay the bills. Only good can come of this.**

以领域概念命名可以改善你和顾客之间的沟通, 这只能带来好处。

+ When you’re struggling to find a good name but have only a few concrete instances to guide you, it can be illuminating(富于启发性的) to imagine other things that would also be in the same category.

当你在为命名苦恼的时候, 可以让一些具体的实例指引你

想象其他同样属于同一范畴的事物也会很有启发性

+ For example, if the song were about wine, the wine might come in a carafe. Juice sometimes comes in small boxes. Soft drinks come in cans.

例如: 如果这首歌是关于"葡萄酒"的, 葡萄酒是carafe(一种饮料瓶)装的; 果汁是盒装的; 软饮料(???)是罐装的

+ If you were to ask your users, "What kind of thing is a bottle?," they wouldn’t reply "It’s a unit."

如果你问其他人, "瓶子是什么东西？", 人们不会回答"是单位"

+ Instead they might call it the container. In the context of 99 Bottles, container is a good name for this concept. Container is meaningful, understandable, and unambiguous.

人们更可能回到 "瓶子是容器(container)"

在这个领域里, **container是一个好名字**

因为它有意义, 易于理解, 领域相关, 毫不含糊

## 3.7.4. Making Methodical Transformations

进行有条不紊的转变

+ Now that you’ve decided to create a container method, it’s time to alter the code.

既然决定了要定义一个`container`方法, 准备开始修改代码

+ It’s tempting to make all of the necessary changes in one fell swoop.

这时候, 倾向于一下做出所有的改变时很有诱惑力(正常)的想法(思路)

+ Doing so requires adding a new method and invoking it in two places.

但是这样需要做两件事: 1. 增加新方法 2. 将两处字符串替换为该方法

*步子太大了*, 因为如果一下做这么多改变, 会增加很多行的代码

+ But wait. Notice that the above change adds seven new lines of code, changes two existing ones, and alters code in three separate places.

上面提到的改动涉及很多行代码

+ Any of these changes could introduce errors, which you would then be obliged to understand and correct.

每行改动都可能引入错误, 你必须理解并纠正这些错误

+ This small example stands in for the much bigger real-life problem where, in the process of implementing a new feature, you add many lines of code, change many others, and then run the tests, only to **be confronted with a ocean of red**.

这个例子扩大了说就是真实的问题, 在实现需求时, 做了很多改动, 增加了很多行代码, 然后运行测试, 发现一片红色的测试不通过的提醒

+ Real world problems are big. Real code has bugs. Real tests are often tightly coupled to current implementations.

真实世界的问题会很大(复杂), 真实世界的代码会有bug, 真实世界的测试常常紧密地和具体实现耦合

+ If you simultaneously change many things and something breaks, you’re forced to understand everything in order to fix anything.

如果同事修改很多处代码, 然后有地方测试不通过, 为了修复这个问题, 你就必须一下子完全掌握所有这些改动(mind burden)

+ You could end up chasing after red, with increasing desperation, before eventually discarding all of the changes and beginning anew.

最终很可能会越来越绝望地debug, 修复测试不通过的问题(修了一个, 另一个又冒出来...), 最终放弃所有的改动, 从头开始

+ Making a slew of simultaneous changes is not refactoring—it’s rehacktoring.

同时进行大量更改不是重构, 而是重新打包

+ It would be much better to make a series of tiny changes and run the tests after each.

最好是进行一系列微小的更改, 并在每次更改后运行测试

+ If the tests fail, you know the exact change that caused the failure, and can undo back to green and make a better change.

如果测试没通过, 可以很容易发现是哪个改动导致的; 定位和处理起来会很简单

+ If the tests pass, you know that the current code works, even if the refactoring is only partially complete.

如果测试通过了, 可以放心的继续进行进一步重构

+ Formal refactoring confers two additional benefits.

正规的重构带来了另外两个好处

+ First, because no change breaks the tests, the code can be deployed to production at any intermediate point. This allows you to avoid accumulating a large set of changes and suffering through a painful merge.

首先是 因为改动都没有让测试失败, 这部分改动的代码能够在任何合适的时间部署到生产环境

这个可以避免积累大量代码改动，避免合并代码带来的痛苦

+ Next, code that runs properly even in the midst of a long refactoring increases the bus factor. This contributes to a higher likelihood of project success even if you, personally, were to meet an untimely end.

其次是 即使在长时间的重构过程中也能正常运行的代码会增加`bus factor`, 这有助于提高项目成功的可能性

> The `bus factor` is a measurement of the risk resulting from information and capabilities not being shared among team members, derived from the phrase "in case they get hit by a bus".

+ As a reminder, here again are the Flocking Rules and corollaries(推论):

作为提醒，这里再次列出了`Flocking Rules`和其推论：

1. Select the things that are most alike.
2. Find the smallest difference between them.
3. Make the simplest change to remove that difference:
    a. parse the new code
    b. parse and execute it
    c. parse, execute and use its result
    d. delete unused code

+ As you’re following the rules:
    + As you’re following the rules:
    + In general, change only one line at a time. Run the tests after every change.
    + If you go red, undo and make a better change.

```bash
cd $HOME/OuterGitRepo/1st_99bottles_ruby && git checkout chapter-3 && tig
```








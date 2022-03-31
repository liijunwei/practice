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


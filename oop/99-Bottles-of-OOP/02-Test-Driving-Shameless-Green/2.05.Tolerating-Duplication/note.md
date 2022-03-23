+ The choice of the best alternative for verse 3 was guided both by metrics and the "Transformation Priority Premise," and those things might again be useful here.
(page 57 Table 2.2)

+ As was stated in the previous section, as tests get more specific, code should become more generic.

前面说过, 随着测试用例变得更加具体, 代码应该变得更加通用

+ Code becomes more generic by becoming more abstract.

代码通过变得更抽象 以获得 通用性

+ One way to make code more abstract is to DRY it out, that is, to extract duplicate bits of code into a single method, to give that method a name, and then to refer to the code by this new name.

一种方法是通过DRY, 把重复的代码提取为方法, 在各处调用方法

+ DRYing out code removes the duplication and thus reduces its overall size.

+ The fact that "bottle" is duplicated many times signals that there’s an underlying concept that has not yet been unearthed.

"bottle" 被重复了很多次, 说明有隐藏的还未被挖掘出的概念

+ DRY is important but if applied too early, and with too much vigor, it can do more harm than good. When faced with a situation like this, ask these questions:
    + Does the change I’m contemplating make the code harder to understand?
    这个改动让代码变得更难理解了吗? 当抽象做得好的时候, 代码会变得更容易理解; 如果某个抽象让代码变得更难懂, 那么说明 做出这个改动的人还没有对问题有清晰的理解
    + What is the future cost of doing nothing now?
    如果现在什么也不做, 以后再做这件事的代价是什么?
    有些改动早做完做的代价差不多, 拖延做这个改动不会使得成本升高; 有时候延迟做出调整, 收集更多的信息或许是更好的方法, 甚至不做改变, 还能省不少钱
    + When will the future arrive, i.e. how soon will I get more information?
    需要多长时间才能看到变化, 什么时候才能收集到更多信息?
    暂时容忍代码的冗余是可接受的
    容忍一定程度的冗余 好过 做出错误的抽象



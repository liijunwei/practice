# 9.1. Choosing Which Units to Test

+ It’s important to know if something breaks, but tests can do far more.
    + They give you the opportunity to explain the domain to future readers.
    + They expose design problems that make code hard to reuse.
    + Well-designed code is easy to test; when testing is hard, the code’s design needs attention.

`测试`能告诉我们 哪里出了问题, 但是`测试`能做的事远比这个多

它们让代码未来的维护者更容易理解代码是如何工作的

它们能暴露出难以重用的代码的设计问题

它们能告诉我们某部分代码是否容易测试, 如果很难测试, 说明代码的设计可能有问题, 需要关注

+ Writing tests, even at this late date, will improve the code. As a final exercise, this chapter does just that.

即使是在很晚的时候才编写测试用例(即 不是TDD), 也能提升代码质量

本章要做的事就是为我们的应用程序补充测试用例

+ Every class should have its own unit test, unless doing otherwise saves money.

每个类都需要有他自己的单元测试, 除非不这么做的时候能省很多钱...

+ The allowed-to-skip-tests bar is high, but some code meets it. This first section explores such a case.

有些部分可以不测试, 但是满足这个条件的门槛很高很高, 有些代码有这样的特征, 我们首先要讨论的就是这种情况

## 9.1.1. Contrasting(对比) Unit and Integration Tests

+ Improving the tests will accomplish two things
    + First, this will reduce costs.
        + Tests that tell the right story make it easier for future readers to understand the code.
        + **Improving the story will save you money forever.**
    + Next, updating the tests will lay bare the consequences of the code’s design.
        + It should be easy to create simple, intention-revealing tests.
        + When it’s not, the chief problem is often too much coupling.
        + In such cases the solution is not to write complicated tests that overcome tight coupling, but rather to loosen the coupling so that you can write simple tests.
        + The most cost- effective time to intervene in tightly coupled code is right now, before new requirements cause you to to start reusing these objects.

提升测试代码质量能完成两件事

1. 降低软件的总成本
    + 测试用例应该能帮助理解软件; 提升测试代码的质量, 能使得用例代码"表达出"一个更加符合真实情况的"故事"
2. 更新测试将揭示代码设计的后果
    + 如果设计得当, 测试用例的编写应该是简单的, 能直接看出意图的
    + 如果没法写出简单的用例, 那么很可能代码里包含了太多的耦合
    + 这种情况下, 不应该是写出复杂的,适应当前代码的测试用例, 去测试这种有很多耦合的代码, 而应该想办法降低现有代码里的耦合, 从而使得写测试变得容易
    + 干预紧耦合代码最具成本效益的时间是现在, 在新需求导致你开始重用这些对象之前(越早越容易)

+ The good news is that writing tests will uncover every bit of overlooked tight coupling and immediately reward you for fixing it.

写测试用例的过程中, 你能发现

好消息是，编写测试将发现所有被忽视的紧密耦合，并立即奖励您修复它。








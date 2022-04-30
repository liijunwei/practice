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








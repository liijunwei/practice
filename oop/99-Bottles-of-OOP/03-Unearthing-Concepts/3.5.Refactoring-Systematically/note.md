+ Having bandied the word around repeatedly, it’s high time for a formal definition of "refactoring." According to Fowler:

在反复讨论"重构"这个词后, 是时候正式解释"重构"这个词的定义了, 据*Martin Fowler*描述:

> Refactoring is the process of changing a software system in such a way that it does not alter the external behavior of the code yet improves its internal structure.

    + 重构是一个过程

    + 这个过程会修改软件系统

    + 以一种修改代码, 但同时不修改代码外部行为的方式 来优化代码的内部结构

重构是一种 以修改代码但同时不修改代码的外部行为的方式优化代码内部结构以有优化整个软件系统的过程

+ In short, refactoring alters the arrangement of code without changing its behavior.

简单来说: 重构会修改代码的排列, 但是不修改代码的行为

+ Recall that **new requirements should be implemented in two steps**
    + First, you rearrange existing code so that it becomes open to the new requirement.
    + Next, you write new code to meet that requirement.

+ Note that safe refactoring relies upon tests.
+ If you truly are rearranging code without changing behavior, at every step along the way the existing tests should continue to pass.
+ Tests are a safety blanket that justifies confidence in the new arrangement of code.
+ If they begin to fail, one of two things must be true.
    1. you’ve inadvertently broken the code
    2. the existing tests are flawed

要注意: 安全的重构依赖于(自动/手动)测试

重构的过程中, 每一步都需要注意现有代码能让测试通过

测试是一个"safe blanked"，可以证明对重构后代码的信心(???怎么翻译呢)

如果测试不通过了, 那么有两种情况至少其一有问题: 1. 重构后的代码行为改变了 2. 之前写的测试用例有问题(too coupled with code)

+ If tests fail because you’ve broken the code, the cure is simple. Undo the last change, make a better one and proceed merrily along your way.

如果刚刚的改动让测试没通过, 只需要撤销改动, 重新编写就行了(需要注意每次改动要小, 确保能准确定位到是哪里引入了问题)

+ However, if you rearrange code without changing behavior and tests begin to fail, then the tests themselves are flawed.

也需要注意: 如果你确定修改后的代码没有改变原有行为, 但是测试用例却失败了, 那说明测试用例本身有缺陷

+ Tests that make assertions about how things are done, rather than what actually happens, are the prime contributors to this predicament.

**测试用例里面, 断言实现细节(事情如何完成)，而不是断言结果(实际发生的事情), 是造成这种问题的主要原因**

+ When in this situation, there’s no alternative other than to improve the tests before embarking upon a refactoring.

这种情况下, 应该先着眼于修改测试用例, 再去重构



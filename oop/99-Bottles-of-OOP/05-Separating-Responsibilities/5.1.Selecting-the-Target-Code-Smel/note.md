+ Code should be open for extension and closed for modification.

![](../open-closed-flowchart.png)


+  If anything, the current incarnation is less amenable to this requirement than was Shameless Green.

当前代码虽然已经移除了一些code smell, 从代码中提取出了一些隐含的概念, 但是还不是符合开闭原则的代码

+ The truth about refactoring is that it sometimes makes things worse, in which case your efforts serve gallantly to disprove an idea.

关于重构的事实是: 重构有可能使得现状变得更糟

+ The refactoring recipes don’t promise to result in code that better expresses the problem— they merely make it easy to create that new expression, and just as easy to revert it.

重构不能保证最终的代码表达能力更强, 能更好的描述问题

重构仅能使得代码更易于添加代码和回滚代码变更

+ Proper refactoring allows you to explore a problem domain safely.

正确的重构能让你对问题的探索变得容易和安全

+ You’ve now completed one refactoring, and the resulting code is not yet open to the six-pack requirement.

+ The current code, although not open to the new requirement, is improved.

当前代码虽然还没有符合开闭原则, 但是毫无疑问代码质量已经有了提升

+ Therefore, have faith, and iterate. This means you must continue to be guided by code smells, and doing so requires that you identify the smells in the current code.

所以打起精神来, 对代码进行迭代吧

这意味着你必须持续的消除代码里的code smell, 通过这种方式优化代码, 迭代地对代码进行重构

## 5.1.1. Identifying Patterns in Code

+ One way to get better at identifying smells is to practice describing the characteristics of code.

一个练习识别code smell的方法是: 多去练习描述某段代码的特征

+ The following questions draw attention to a number of interesting characteristics of the code as it’s written so far:

这几个问题能引发你对代码几个有趣的特点的思考

1. Do any method shave the same shape ?
2. Do any methods take anargument of the same name ?
3. Do arguments of the same name always mean the same thing ?
4. If you were to add the private keyword to this class, where would it go ?
5. If you were going to break this class into two pieces, where’s the dividing line ?







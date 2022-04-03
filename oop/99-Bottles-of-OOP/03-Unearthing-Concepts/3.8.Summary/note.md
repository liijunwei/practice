+ When faced with the need to change code, very often the hardest decision is where to start.

当需要修改代码时, 通常来说最困难的选择是"从哪里开始"

+ This chapter suggested that you be guided by the Open-Closed Principle, and so separate most changes into two broad steps.

本章节建议遵循"开闭原则"来进行代码改动

这些改动大概分为两大步:

+ First, refactor the existing code to be open to the new requirement.

首先, 把现有代码重构为对新需求的扩展开放的状态

+ Next, add the new code.

然后, 添加新代码, 扩展新功能

+ Sometimes the first step, refactoring to openness, requires such a large leap that it is not obvious how to achieve it.

很多时候, "把现有代码重构为对新需求的扩展开放的状态" 需要做很多工作, 以至于目标过大, 而不知道该从那里开始(常见问题)

+ In that case, be guided by code smells.

这种情况下, 应该先消除"Code Smell"

+ Improve code by identifying and removing smells and have faith that as the code improves, a path to openness will appear.

通过识别并消除"Code Smell"的方法来改进代码, 并且随着代码的改进, 一条通向"把现有代码重构为对新需求的扩展开放的状态"的道路会渐渐浮现

+ Making existing code open to a new requirement often requires identifying and naming abstractions.

"把现有代码重构为对新需求的扩展开放的状态"通常需要识别并为做出的抽象命名

+ The Flocking Rules concentrate on turning difference into sameness, and thus are useful tools for unearthing abstractions.

[`Flocking Rules`和其推论](../3.7.Converging-on-Abstractions/note.md#374-making-methodical-tran)专注于将差异转化为相同, 因此是挖掘抽象的有用工具

+ However, now that you’ve learned how to use the flocking rules to identify abstractions, resolving the differences in the 1 and 0 cases will go much faster.

+ So, on to Chapter 4, and more extracting of abstractions.

第四章里还有更多`抽象`相关的内容




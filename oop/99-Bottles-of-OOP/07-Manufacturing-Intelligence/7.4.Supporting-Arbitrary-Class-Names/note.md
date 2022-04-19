+ The first step towards an open factory that both centralizes knowledge and supports arbitrary class names is to rearrange the code to increase the isolation of the names.

想要让工厂 既能将知识集中管理, 又让它对任意名字的扩展开放, 第一步需要做的是重新组织代码, 提升这些名字的隔离程度

+ You can do this by replacing the case statement with a key/value lookup, as follows:

可以把switch-case替换为键值对查找来实现



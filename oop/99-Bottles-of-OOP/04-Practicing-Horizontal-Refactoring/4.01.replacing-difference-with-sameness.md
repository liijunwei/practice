+ This chapter iteratively applies the Flocking Rules to the remaining special verses, and results in a single, more abstract, template that produces every possible verse.

本章将继续一步一步应用"the Flocking Rules" 重构剩余的代码

然后会生成一个更抽象的模板, 生成所有可能的诗句(verse)

+ The refactoring rules say to start by choosing the cases that are most alike.

+ Of these three verse templates, the 1 and else cases are most alike, so
they’re the next to address.

+ Replacing differing concrete values with a reference to a common variable changes difference into sameness.

将具体实例里的不同之处替换为对公共变量的同一引用能将不同的两个分支变成一样的

+ This substitution is important, not because it changes the resulting value, but because it increases the level of abstraction.

这种替换很重要, 不是因为他改变了值, 而是因为它做出了一层抽象

+ It is this increase in abstraction that makes things the same. Without it, you are doomed to the conditional.

正是这层抽象, 使得两个分支变得更加相似, 然后合并为同一分支

如果没有它, 那么必然需要使用条件判断

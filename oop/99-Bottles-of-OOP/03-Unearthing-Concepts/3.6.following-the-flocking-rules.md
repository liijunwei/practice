+ Considered from a higher viewpoint, each variant is merely a verse in the song; in that sense they are all the same.

从较高的维度去思考, 每个case, 都是歌词的一部分, 从这个角度看, 他们都是一样的

+ Underlying each concrete variant is a generalized verse abstraction.

每个case都是一种潜在/通用抽象的变种

+ If you could find this abstraction, you could use it to reduce the four-branch case statement to a single line of code.

如果你能把他们抽象出来, 你就能用它把多个case合并为一个case

+ The good news is that you don’t have to be able to see the abstraction in advance.

好在你不需要一下子就能识别出合适的抽象层次

+ You can find it by iteratively applying a small set of simple rules.

你可以通过不断地/迭代地应用一些法则, 一点点发现潜在的抽象

+ These rules are known as "Flocking Rules", and are as follows:

这些法则叫做"Flocking Rules", 他们的表述如下

1. Select the things that are most alike.

选出代码中最相似的部分

2. Find the smallest difference between them.

选出他们中最不同的部分

3. Make the simplest change that will remove that difference.

做一些微小的调整, 以消除这些不同

+ Changes to code can be subdivided into four distinct steps:

对代码的改动可被分为下面几个清晰的步骤(??? TODO)

1. parse the new code
2. parse and execute it
3. parse ,execute and use its result
4. delete unused code

+ Making small changes means you get very precise error messages when something goes wrong, so it’s useful to know how to work at this level of granularity.

一点点调整意味着如果有问题, 能及时发现那里出了问题

在这个粒度上工作很有用(??), 开始时步子要小, 要踏实

+ As you gain experience, you’ll begin to take larger steps, but if you take a big step and encounter an error, you should revert the change and make a smaller one.

当逐渐熟练以后, 就可以提升一些速度, 粒度可以加大

但是在你步子变大后, 遇到问题时, 要及时回滚改动, 然后用更小的粒度去做重构

+ As you’re following the flocking rules:
    + For now, change only one line at a time.
    + Run the tests after every change.
    + If the tests fail, undo and make a better change.

### Why "Flocking"

+ The group’s behavior is the result of a continuous series of small decisions being made by each participating individual.

群体的行为是 参与者一系列连续的小的选择 的结果

+ These decisions are guided by three simple rules.

这些选择由下面这三条简单的规则决定

1. Alignment-Steer towards the average heading of neighbors

对齐: 转向邻居的平均航向(?)

2. Separation-Don’t get too close to a neighbor

分隔: 不要举例过近

3. Cohesion-Steer towards the average position of the flock

凝聚: 转向平均位置

+ Thus, complex behavior emerges from the repeated application of simple rules.

复杂的行为由一系列简单/重复的原则的应用组合而来

+ In the same way that the rules in this sidebar allow birds to flock, the "Flocking Rules" for code allow abstractions to appear.

(???)

+ To see a beautiful example of flocking in action, watch [Steven Strogatz’s The Science of Sync TED talk](https://www.youtube.com/watch?t=196&v=IiXaZGZqpVI&feature=youtu.be).


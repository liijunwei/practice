You have two methods in subclasses that perform similar steps in the same order, yet the steps are different.

*Get the steps into methods with the same signature, so that the original methods become the same. Then you can pull them up.*

+ 找出相同的部分, 抽象出一个方法, 这个方法就是 template
    + 利用 继承
    + 或者
    + 利用 block
    + 把不一样的部分 用多态(and super?) 或 block 给实现了...

+ The key to this refactoring is to **separate the varying code from the similar code by using `Extract Method` to extract the pieces that are different between the two methods**.

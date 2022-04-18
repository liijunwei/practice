+ This selection happens in the `BottleNumber.for` factory.

上一章节中最核心的工厂方法是`BottleNumber.for`

+ In this chapter, that simple factory launches a greater exploration of factories in general.

本章中, 将对工厂进行更加深入的探索

+ Given that these different conditionals produce the same variability, can you explain why one contains four branches, but the other only three?

对比"shameless green"和`BottleNumber.for`, 你能解释为什么他们机构相似, 但是前者有4个分支, 后者只有3个分支吗?

+ To answer that question, consider the case that is missing.
+ **Shameless Green has a special case for 2, but the factory does not.**

+ Recall that the conditional in Shameless Green produces verses, but the one in the factory produces bottle numbers.

+ Verse 2 is indeed special, but bottle number 2 is not.

verse2是一种特殊情况, 但是 bottle_number2 不是特殊情况

+ Thus, Shameless Green needs a special case for verse 2 solely because verse 2 contains bottle number 1.
+ This explains the missing branch.

前者(Shameless Green)里的条件判断知道这些事:

1. 判断和选择的理由
2. 每个分支里的具体行为

后者(`BottleNumber.for`)里的条件判断知道这些事:

1. 判断和选择的理由(相似)
2. 只知道 每个分支里的BottleNumber, 不知道具体行为(主要的不同)

+ Factories don’t know what to do: instead, they know how to choose who does.

工厂不知道该做什么, 它只知道如何选择谁来做什么

+ They consolidate the choosing and separate the chosen.

他们把做选择和选择结果的具体后果(行为)区别开了

+ Shameless Green was a procedure because it combined these two things;

"Shameless Green"是面向过程的, 因为他把选择和结果绑定在了一起

+ The current code is object-oriented because it breaks them apart.

但是当前版本的代码是面向对象的, OOP里最好把选择和行为区分开(面向接口编程)

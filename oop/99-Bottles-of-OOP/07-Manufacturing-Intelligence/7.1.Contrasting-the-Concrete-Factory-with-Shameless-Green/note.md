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


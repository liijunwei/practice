**This chapter explored the Data Clump code smell. It replaced a Switch Statement with a set of polymorphic objects, which it created using a factory.**

本章探索了"Data Clump"的code smell, 并且使用多多态+工厂方法将条件判断消除掉了

The BottleNumber for factory was straightforward and most certainly did the job.

BottleNumber的`for`工厂方法很直白, 并且很好用

+ While simple factories like this work great in many situations, they’re not best for every case. There’s a whole world of different styles of factories waiting to be explored.

虽然在多数情况下这种工厂都很好用, 但是在有些情况下他们并不适用

还有很多很多不同类型的工厂可以探索和使用


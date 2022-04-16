+ In OO, polymorphism refers to the idea of having many different kinds of objects that respond to the same message.

在面向对象编程里, 多态指的是有很多对象都实现了名称相同的方法, 但是各自的行为不同

+ TODO 问题: sender 和 receiver 怎么区分... 迷糊了

## 6.3.1. Dismembering Conditionals

+ Polymorphism, by definition, involves more than one kind of object, so changing from a procedural to a polymorphic code arrangement will increase the overall number of classes.

多态 从定义上来看就是和多个种类的对象有关, 所以在将条件判断的代码转变为过程式代码时, 一定会增加总的类的数量

+ The specific logic for 0 needs to be isolated in a class of its own, as does the logic for 1.
+ Also, as these new classes come into existence, some additional code will have to be written to choose the correct class based on the value of number.


## 6.3.2. Manufacturing Objects

+ When several classes play a common role, some code, somewhere, must know how to choose the right role-playing class for any specific contingency.

当有多个类扮演同一个角色时, 必须有部分代码知道该选哪个类来完成这件事

+ This choosing very often involves a conditional, which should exist in one and only one place.

这种选择通常包含了条件判断, 但是这种判断应该只存在一处, 并且只存在一处

+ Code like this is said to "manufacture" an instance of the right kind of object, and so is commonly referred to as a factory.

这种做出选择(决定类型的代码)的过程, 叫做"制造出"对应的实例, 所以也被叫做"工厂", 第7章将详细介绍工厂


## 6.3.3. Prevailing with Polymorphism


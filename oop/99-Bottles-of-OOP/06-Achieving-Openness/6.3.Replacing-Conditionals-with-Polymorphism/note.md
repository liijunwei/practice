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

+ When a factory exists for a role, the factory has sole responsibility for creating objects to play that role.

当存在工厂的时候, 它唯一的职责就是为使用它的人创建合适的对象

+ The factory’s purpose is to isolate the names of the concrete classes, and to hide the logic needed to choose the correct one.

工厂的目标就是把具体的类名隔离开, 隐藏选择的逻辑

+ Now that BottleNumber0 exists, you need a bottle number factory.

+ The first step is to do a small refactoring to isolate the creation of bottle numbers in a single method of Bottles.

+ These classes are substitutable for one another. When you invoke the factory to get a bottle number, you don’t need to know the class of the returned object.

这些类之间可以相互替换, 当你使用工厂方法获取一个实例的时候, 你不需要关注工厂返回了哪个类的实例

+ You merely trust that object to act like a bottle number and to respond to the messages you plan to send.

你只需要信任这个实例的行为是一个瓶子的数量, 并且对你发出的消息有响应就好了

## 6.3.3. Prevailing with Polymorphism

Here’s a list of the recipe’s steps:

1. Create a subclass to stand in for the value upon which you switch.
    + a. Copy one method that switches on that value into the subclass.
    + b. In the subclass, remove everything but the true branch of the conditional.
        + i. At this point, create a factory if it does not yet exist, and
        + ii. Add this subclass to the factory if not yet included.
    + c. In the superclass, remove everything but the false branch of the conditional.
    + d. Repeat steps a-c until all methods that switch on the value are dispersed.
2. Iterate until a subclass exists for every different value upon which you switch.

+ The conditional above may be giving you "a sense of deja vu".

上述条件可能会给你一种似曾相识的感觉

+ It’s reminiscent of, although not quite identical to, the case statement from the original Shameless Green solution.

这个case语句, 和"shameless green"里的分支很相似

+ Think about why this might be as you finish the current refactoring. The similarity will be explored at the end of this section.

思考一下为什么会有相似的感觉

这种相似性会在本节结束时再讨论






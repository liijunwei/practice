+ Each item above acts like a vote, and these votes combine to point to Primitive Obsession as the dominant code smell.

前面每个问题的答案都像一张选票, 这些选票结合起来表明 对"基本数据类型"的痴迷

+ Built-in data classes like String , Fixnum , Integer ( Fixnum 's superclass), Array , and Hash are examples of "primitives."

String Fixnum Integer Array Hash 等 就是 "基本数据类型"

+ **Primitive Obsession is when you use one of these data classes to represent a concept in your domain.**

"对原始类型的痴迷"意思是: 使用了那些基本数据类型代表你领域问题里的一个概念

+ Obsessing on a primitive results in code that passes built-in types around, and supplies behavior for them.

其结果就是: 代码里把基本数据类型传来传去, 并且为这些基本数据类型赋予行为

(should pass object around instead of pass primitive data type around)

+ The cure for Primitive Obsession is to create a new class to use in place of the primitive. For this operation, the refactoring recipe is Extract Class.

"对原始类型的痴迷"的解药就是 由传递基本数据类型 改为 传递类的实例

## 5.2.1. Modeling Abstractions

+ This new class does not represent a kind of bottle: it represents a kind of number.

用来替代"原始数据类型"的类, 不是一类"瓶子", 而是一类数字

+ A bottle is a thing, while a number is an idea. It’s easy to imagine creating objects that stand in for things, but the power of OO is that it lets you model ideas.

"瓶子"是一类事物, 而数字是一类概念

OOP的好处之一是, 它能让你为概念抽象出对应的类

+ Model-able ideas often lie dormant in the interactions between other objects.

可建模的想法通常隐藏在其他对象之间的交互中

+ Embodying these concepts into discrete classes separates responsibilities and makes the overall application easier to understand, test, and change.

+ The two most obvious choices are BottleNumber , or ContainerNumber.

## 5.2.2. Naming Classes

+ BottleNumber is more concrete. ContainerNumber is more abstract.

+ The tie-breaker here is that the "name things at one higher level of abstraction" rule applies more to methods than to classes.

"name things at one higher level of abstraction" 这个规则, 和类(class)相比, 它更适合于方法(method)

+ The rule about naming can thus be amended: while you should continue to name methods after what they mean, classes can be named after what they are.

+ As always, you can revisit this decision if things change later.

## 5.2.3. Extracting BottleNumber

+ This section extracts a new class named BottleNumber from the existing code.

the process of changing code was subdivided into four steps.
1. parse the new code
2. parse and execute it
3. parse, execute and use its result
4. delete unused code

+ Start the class extraction by creating an empty BottleNumber class,asshownbelow:

+ Remember that the verse method should not be extracted. Even though its argument is also named number , in this case the argument represents a verse number, not a bottle number.

verse方法里的number和其他方法里的number是不同的概念

+ As mentioned earlier, the recipe being followed here was inspired by one from Martin Fowler. The "official" Extract Class recipe begins by linking the old class to the new. Then, one at a time, the recipe moves fields, and then methods, of interest. In contrast, the example above starts with Fowlers final step, and combines all of the method moves within a single change.

提取类的思路是: 每次提取一小部分, 然后运行测试, 保证每次的改动都是测试通过的

+ First, fully connect BottleNumber to Bottles . Once that’s complete, you can return and improve the methods in BottleNumber .

先用 BottleNumber 类把Bottle里的代码替换掉

+ Repeat the above procedure for each of the methods copied from the Bottles class. This is an extremely mechanical, wonderfully boring, and deeply comforting refactoring process.

这种重构方式, 安全/简单/容易/舒服...

## 5.2.4. Removing Arguments

+ The point of the Primitive Obsession/Extract Class refactoring is to create a smarter object to stand in for the primitive.

使用类替代原始数据类型的要点就是 使用它们(class/object)代替原始数据类型发送消息

+ This smarter object, by definition, knows both the value of the primitive and its associated behavior.

这些"smarter objects"知道基本数据类型对应的值和他们的行为

+ Because the new BottleNumber class holds the right number, the methods in BottleNumber don’tneedtotakeanargument,andinvokersofthesemethods could be relieved of their obligation to pass a parameter.

+ Keep in mind that is a multi-line change. Some problems are so simple that it’s easiest just leap in and make such a change, but others are so complex that it isn’t feasible to fix everything at once.

要提醒自己, 每次做少量的改动, 然后运行测试, 是一种"最佳实践", 需要多多练习, 形成这种习惯

## 5.2.5. Trusting the Process

+  If you adhere to a recipe and tests start failing, it’s likely that there’s something about the problem that you don’t yet understand.

SO TRUE

+ Recall the steps needed to remove parameters:

1. Alter the method definition to change the argument name, and provide a default.
2. Change every sender of  the message toremove the parameter.
3. Delete the argument from the method definition.

**The failure appeared after step 3.**

The error message indicates that some caller is still passing a parameter to pronoun.

This means step 2 isn’t complete; i.e. some sender has not been fixed.

+ The lesson here is that the process works, and that encountering errors while following it suggests that a closer look at the code is in order. A great benefit of these refactoring techniques is that you can accomplish quite a bit while thinking very little.





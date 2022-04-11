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






















®
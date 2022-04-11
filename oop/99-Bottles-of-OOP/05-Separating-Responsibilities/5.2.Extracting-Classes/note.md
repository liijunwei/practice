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

+ A bottle is a thing, while a number is an idea. It’s easy to imagine creating objects that stand in for things, but the power of OO is that it lets you model ideas.



































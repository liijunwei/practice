+ Every piece of knowledge is a dependency, and the way that quantity is written requires verse to know too many things.

所有必须的知识都是 依赖

`quantity`方法的实现方式, 使得verses方法需要知道太多的东西了

+ The idea of reducing the number of dependencies imposed upon message senders by requiring that receivers return trustworthy objects is a generalization of the Liskov Substitution Principle.

通过要求接收者返回可信对象来减少对消息发送者施加的依赖的数量, 是Liskov替换原理的一个推广

+ The official definition of Liskov says that "subtypes must be substitutable for their supertypes."

里氏替换原则的定义是: 子类型必须能替换他们的父类型

+ Liskov, in plain terms, requires that objects be what they promise they are.

+ When using inheritance, you must be able to freely substitute an instance of a subclass for an instance of its superclass.


使用继承时, 必须保证能够让子类实例替换其父类的实例

+ Subclasses, by definition, are all that their superclasses are, plus more, so this substitution should always work.

根据定义, 子类就是它们父类的加上扩展功能, 所以这种替换应该总是有效的

+ The Liskov Substitution Principle also applies to duck types.

里氏替换原则也适用于鸭子类型

+ When relying on duck types, every object that asserts that it plays the duck’s role must completely implement the duck’s API. Duck types should be substitutable for one another.

当依赖于鸭子类型时, 每个实例都应该断定那个充当鸭子类型角色的实例已经实现了对应的鸭子类型API

每种鸭子类型的API实现都应该能替换另一种


+ The decision about whether to refactor in the first place should be determined by whether your code is already "open" to the new requirement.

决定是否需要重构代码前, 需要先确定你的代码是否支持 对新需求进行扩展

+ "Open" is short for "Open/Closed," which in turn is short for "open for extension and closed for modification."

Open 是 开闭原则的简写, 表示代码要对扩展开放, 对修改关闭

+ Code is open to a new requirement when you can meet that new requirement without changing existing code.

即: 代码写好以后, 为适应新的需求, 应该尽量少去考虑修改原有代码, 而应考虑去扩展原有代码

## SOLID

+ S - Single Responsibility

The methods in a class should be cohesive(有凝聚力的) around a single purpose.

一个类中的方法应该只做一件事, 并把它做好

+ O - Open-Closed

Objects should be open for extension, but closed for modification.

类的实例应该对扩展开放, 对修改关闭

+ L - Liskov Substitution

Subclasses should be substitutable for their superclasses.

子类应该能替代它的父类

+ I - Interface Segregation

Objects should not be forced to depend on methods they don’t use.

接口隔离原则, 类的实例不应该依赖那些它不使用的方法

+ D - Dependency Inversion

Depend on abstractions, not on concretions.

依赖倒置原则: 程序应该依赖于抽象接口, 不应该依赖于具体实现方式

A.高层次的模块不应该依赖于低层次的模块, 他们都应该依赖于抽象
B.抽象不应该依赖于具体, 具体应该依赖于抽象




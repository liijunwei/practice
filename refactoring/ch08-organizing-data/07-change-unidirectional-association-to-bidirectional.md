You have two classes that need to use each other's features, but there is only a one-way link.

两个类里 都需要用到对方的功能, 但是只有一方有另一方的实例引用(单向的引用)

*Add back pointers, and change modifiers to update both sets.*

+ you need to set up a `two-way reference`, sometimes called a `back pointer`.

+ If you aren't used to back pointers, it's easy to become tangled up using them. Once you get used to the idiom(习语/成语/惯用于法), however, it is not too complicated.

+ The idiom is awkward enough that you should have tests, at least until you are comfortable with the idiom.

+ This refactoring uses back pointers to implement bidirectionality.

+ How to decide which class will control the association
    + If both objects are reference objects and the association is **one to many**, then the object that has the one reference is the controller. (That is, if one customer has many orders, the order controls the association.)
    + 一对多关系中的 1 对应的类是控制方(unclear)

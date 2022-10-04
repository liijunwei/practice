You have one class doing work that should be done by two.

*Create a new class and move the relevant fields and methods(behavior) from the old class into the new class.*

+ A class should be a crisp abstraction, handle a few clear responsibilities, or some similar guideline.
    + 一个类应该是一个清晰的抽象, 用来处理明确的职责
    + In practice, classes grow. You add some operations here, a bit of data there.
    + 但实践中, 类会随着业务增长, 如果不及时明晰职责, 任由它肆意增长, 类会变得过于复杂, 难以理解和修改

+ You need to consider where it can be split, and you split it.
    + A good sign is that a subset of the data and a subset of the methods seem to **go together**.
    + Other good signs are subsets of data that usually **change together** or are particularly **dependent on each other**.
    + A useful test is to ask yourself **what would happen if you removed a piece of data or a method**. What other fields and methods would become nonsense?

+ `two-way link` -> `one-way link`

+ Extract Class is a common technique for **improving the liveness of a concurrent program** because it allows you to have separate locks on the two resulting classes.

+ don't understand
    + "Transactions are useful when you use them, but writing transaction managers is more than most programmers should attempt."

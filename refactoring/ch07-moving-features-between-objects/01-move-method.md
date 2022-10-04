+ TODO read again(complex)

+ Moving methods is the bread and butter of refactoring(重构的基础)

+ I move methods when classes have too much behavior or when classes are collaborating too much and are too highly coupled
+ By moving methods around, I can make the classes simpler, and they end up being a more crisp implementation of a set of responsibilities(一组更清晰的实现)

+ It’s not always an easy decision to make.
    + If I am not sure whether to move a method, I go on to look at other methods.
    + Moving other methods often makes the decision easier.
    + Sometimes the decision still is hard to make.
    + Actually it is then no big deal. If it is difficult to make the decision, it probably does not matter that much.
    + Then I choose according to instinct; after all, **I can always change it again later.**

+ Sometimes it is easier to move a clutch of methods than to move them one at a time.

+ If the method includes exception handlers, **decide which class should logically handle the exception**. If the source class should be responsible, leave the handlers behind.

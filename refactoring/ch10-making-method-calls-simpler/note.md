**Objects are all about interfaces.**

Coming up with interfaces that are easy to understand and use is a key skill in developing good Object-Oriented software.

This chapter explores refactorings that **make interfaces more straightforward**.

**Often the simplest and most important thing you can do is to change the name of a method.**

Naming is a key tool in communication.

If you understand what a program is doing, you should not be afraid to use `Rename Method` to pass on that knowledge.

You can (and should) also rename variables and classes.

On the whole these renamings are fairly simple text replacements, so I haven't added extra refactorings for them.

+ Usually you can replace long parameter lists with **immutable objects**, but otherwise you need to be cautious about this group of refactorings.
    + 如果数据是可变的, 在并发情况下可能(一定)会出问题

+ One of the most valuable conventions I've used over the years is to clearly **separate methods that change state (modifiers) from those that query state (queries).**
    + `Separate Query from Modifier`

+ Good interfaces show only what they have to and no more.
    + hide/remove unnecessary methods
    + hide data
    + (immutable when possible?)
    + simple/clear interface

+ Constructors are a particularly awkward feature of Ruby and Java, because they force you to know the class of an object you need to create.
    + Often you don't need to know this.
    + The need to know can be removed with `Replace Constructor with Factory Method`.

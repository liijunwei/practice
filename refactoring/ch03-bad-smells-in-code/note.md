+ Divergent Change 发散的变化

+ Divergent change occurs when one class is commonly changed in different ways for different reasons.

+ Shotgun surgery is similar to divergent change but is the opposite. You whiff this when every time you make a kind of change, you have to make a lot of little changes to a lot of different classes.

+ Divergent change is one class that suffers many kinds of changes, and shotgun surgery is one change that alters many classes.

+ The whole point of objects is that they are a technique to package data with the processes used on that data.

+ The fundamental rule of thumb is to put things together that change together.

+ Data Clumps
    + Don't worry about data clumps that use only some of the attributes of the new object. As long as you are replacing two or more instance variables with the new object, you'll come out ahead.
    + Reducing instance variable lists and parameter lists will certainly remove a few bad smells, but once you have the objects, you get the opportunity to make a nice perfume.

+ Primitive Obsession(mind blowing)
    + Replace Data Value with Object on individual data values

+ Case Statements
    + The problem with case statements is essentially that of duplication
    + Often you find the same case statement scattered about a program in different places
    + Most times when you see a case statement you should consider polymorphism

+ Lazy Class
    + Each class you create costs money to maintain and understand. A class that isn't doing enough to pay for itself should be eliminated.

+ Speculative(推测) Generality(普遍性)
    + You get it when people say, "Oh, I think we need the ability to do this kind of thing someday" and thus want all sorts of hooks and special cases to handle things that aren't required.

+ we can rarely figure out a design until we’ve mostly built it

+ Refused Bequest(遗赠)

+ what is "Introduce Assertion" ???

+ When you feel the need to write a comment, first try to refactor the code so that any comment becomes superfluous.

+ A good time to use a comment is when you don’t know what to do. In addi- tion to describing what is going on, comments can indicate areas in which you aren’t sure. A comment is a good place to say why you did something. This kind of information helps future modifiers, especially forgetful ones.

+ Repetitive Boilerplate
    + Most code isn’t simple enough to declare in this way(e.g.: attr_reader), but when the purpose of the code can **be captured clearly in a declarative statement**, Introduce Class Annotation can clarify the intention of your code.



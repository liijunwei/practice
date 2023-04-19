+ **Code Should be Obvious**

+ The solution to the obscurity problem is to write code in a way that makes it obvious
+ this chapter discusses some of the *factors* that make code more or less obvious.

+ obvious的代码有这些特点
    + 其他人能快速读懂代码，并且他们对代码的直观猜测就是代码想要表达的意思
    + 其他人能快速获取到代码的关键信息

+ “Obvious” is in the mind of the reader: it’s easier to notice that someone else’s code is nonobvious than to see problems with your own code.
    + 代码是否obvious应该由写代码以外的其他人来判断
    + 旁观者清
    + *If someone reading your code says it’s not obvious, then it’s not obvious, no matter how clear it may seem to you.*
    + By trying to understand what made the code nonobvious, you will learn how to write better code in the future.

+ Things that make code more obvious
    + good names(precise and meaningful names)
    + consistency
    + judicious use of white space(合理/明智的 使用空格/空行)
    + comments(you must put yourself in the position of the reader and figure out what is likely to confuse them, and what information will clear up that confusion.)

+ Things that make code less obvious
    + event-driven programming(Event-driven programming makes it hard to follow the flow of control)
    + generic containers(generic containers result in nonobvious code because the grouped elements have generic names that obscure their meaning)
    + different types for declaration and allocation
        + e.g. The variable is declared as a List, but the actual value is an ArrayList
    ```java
    private List<Message> incomingMessageList;
    // ...
    incomingMessageList = new ArrayList<Message>();
    ```
    + code that violates reader expectations


+ Red Flag: Nonobvious Code
    + If the meaning and behavior of code cannot be understood with a quick reading, it is a red flag.
    + *Often this means that there is important information that is not immediately clear to someone reading the code*.

+ **software should be designed for ease of reading, not ease of writing**

+ To make code obvious, you must ensure that readers always have the information they need to understand it.
+ You can do this in three ways:
    1. The best way is to **reduce the amount of information that is needed**, using design techniques such as abstraction and eliminating special cases.
    2. you can take advantage of information that readers have already acquired in other contexts (for example, by following conventions and conforming to expectations) so readers don’t have to learn new information for your code.
    3. Third, you can present the important information to them in the code, using techniques such as good names and strategic comments.
    减少依赖的信息，减少特殊情况
    利用已有的上下文信息，例如遵循惯例
    通过各种手段将重要信息展示在代码中(好的名字，注释等)



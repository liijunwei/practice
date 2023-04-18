+ **Consistency** is a powerful tool for *reducing the complexity of a system* and *making its behavior more obvious*.

+ Consistency creates cognitive leverage: once you have learned how something is done in one place, you can use that knowledge to immediately understand other places that use the same approach.

+ Consistency reduces mistakes

+ Consistency can be applied at many levels in a system, e.g.:
    + names
    + coding style
    + interfaces
    + design patterns
    + Invariants(不变量)
        + An invariant is a property of a variable or structure that is always true. For example, a data structure storing lines of text might enforce an invariant that each line is terminated by a newline character. Invariants reduce the number of special cases that must be considered in code and make it easier to reason about the code’s behavior.
        + 不变量是变量或结构中始终为真的属性，例如约定使用不可变的数据结构，约定functional-service的写法，约定使用Command-query separation(CQS)...

+ how to ensure consistency
    + Document
    + Enforce: The best way to enforce conventions is to *write a tool that checks for violations*, and make sure that code cannot be committed to the repository unless it passes the checker. Automated checkers work particularly well for low-level syntactic conventions.
    + Code reviews: The more nit-picky that code reviewers are, the more quickly everyone on the team will learn the conventions, and the cleaner the code will be.
    + “When in Rome, do as the Romans do.”: When you see anything that looks like it might possibly be a convention, follow it.
    + Don’t change existing conventions: Resist the urge to “improve” on existing conventions. **Having a “better idea” is not a sufficient excuse to introduce inconsistencies.**
        + 一致性有时甚至大于正确性和准确性
        + Overall, reconsidering established conventions is rarely a good use of developer time.

+ Consistency means not only that similar things should be done in similar ways, but that dissimilar things should be done in different ways.
    + 一致性不仅强调相似的事要用相似的方法来处理，还强调不同的事应该用不同的方法来处理

+ Consistency is another example of the investment mindset.

+ It will take a bit of extra work to ensure consistency:
    + work to decide on conventions
    + work to create automated checkers
    + work to look for similar situations to mimic in new code
    + work in code reviews to educate the team
    +
    + **The return on this investment is that your code will be more obvious**
    + Developers will be able to understand the code’s behavior more quickly and accurately, and this will allow them to work faster, with fewer bugs

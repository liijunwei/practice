You have duplicated behavior in two or more classes.

*Create a new module and move the relevant behavior from the old class into the module, and include the module in the class.*

+ reasons
    + removal of duplication

+ A module should have a single responsibility, just like a class.
+ The methods within a module should be cohesive: **They should make sense as a group**.

+ Too often I've seen modules become "junk-drawers" for behavior.
+ They are created with the noble goal of removing duplication, but over time they become bloated.

+ **A module that is difficult to name without using words like "Helper" or "Assistant" is probably doing too much.**

+ When we're talking about the removal of duplication, use `Extract Class` whenever possible.

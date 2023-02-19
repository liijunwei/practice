### small

+ the first rule of functions is that they should be small

### do one thing

+ "level of abstraction"

+ functions that do one thing cannot be reasonably divided into sections

### one level of abstraction per function

+ **in order to make sure our functions are doing "one thing", we need to make sure that the statements within our function are all at the same level of abstraction**
    + mixing levels of abstraction within a function is always confusing

+ my general rule for `switch` statements is that they can be tolerated if they appear only once, are used to create polymorphic objects, and are hidden behind an inheritance relationship so thay the rest of the system can't see them


### use descriptive names

+ you know you're working on clean code when each routine turns out to be pretty much what you expected

+ don't be afraid to make name long, a long descriptive name is better than a short enigmatic name

+ don't be afraid to spend time choosing a name
    + Instead, you should try several different names and read the code with each in place

+ be consistent in your names

### function arguments

+ number of function arguments
    + ideal: 0
    + next: 1
    + next: 2

### flag arguments

+ passing a boolean into a function is a truly terrible practice
    + it loudly proclainming that this function does more than one thing

### argument objects

+ when a function seems to need more than two or three arguments, it is likely that some of those arguments ought to be wrapped into a class of their own

+ monads 一元的/单子
+ dtads  二元的
+ triads 三元的



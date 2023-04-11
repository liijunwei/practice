+ Selecting names for variables, methods, and other entities is one of the most underrated aspects of software design. Good names are a form of documentation: they make code easier to understand.

+ However, software systems have thousands of variables; choosing good names for all of these will have a significant impact on complexity and manageability.

+ It took six months, but I eventually found and fixed the bug. The problem was actually quite simple (as are most bugs, once you figure them out).

+ Thus, you shouldn’t settle for names that are just “reasonably close”. Take a bit of extra time to choose great names, which are precise, unambiguous, and intuitive. The extra attention will pay for itself quickly, and over time you’ll learn to choose good names quickly.

+ **Thus, the challenge is to find just a few words that capture the most important aspects of the entity.**
    + 取名时要 抓住重点，忽略细节

+ **Names are a form of abstraction: they provide a simplified way of thinking about a more complex underlying entity.**
+ Like other forms of abstraction, the best names are those that focus attention on what is most important about the underlying entity while omitting details that are less important.

+ Names should be precise

+ Red Flag: Vague Name

+ If you find it difficult to come up with a name for a particular variable that is precise, intuitive, and not too long, this is a red flag.
+ It suggests that the variable may not have a clear definition or purpose.
+ When this happens, consider alternative factorings.
+ The process of choosing good names can improve your design by identifying weaknesses.

+ Red Flag: Hard to Pick Name

+ Use names consistently
    + first, always use the common name for the given purpose
    + second, never use the common name for anything other than the given purpose
    + third, make sure that the purpose is narrow enough that all variables with the name have the same behavior

+ If you start getting complaints that your code is cryptic, then you should consider using longer names (a Web search for “go language short names” will identify several such complaints). Similarly, if I start getting complaints that long variable names make my code harder to read, then I’ll consider using shorter ones.

+ Well chosen names help to make code more obvious; when someone encounters the variable for the first time, their first guess about its behavior, made without much thought, will be correct.

+ Developing a skill for naming is also an investment.
+ When you first decide to stop settling for mediocre names, you may find it frustrating and time-consuming to come up with good names.
+ However, as you get more experience you’ll find that it becomes easier;
+ eventually, you’ll get to the point where it takes almost no extra time to choose good names, so you will get the benefits almost for free.



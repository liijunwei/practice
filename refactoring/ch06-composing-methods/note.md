# Composing Methods

A large part of my refactoring is composing methods to package code prop- erly.

The biggest problem with Extract Method is dealing with local variables, and temps are one of the main sources of this issue.

+ TODO question: what is "Split Temporary Variable"?

+ When a **default parameter** becomes unused, I need to remove it using Remove Unused Default Parameter.


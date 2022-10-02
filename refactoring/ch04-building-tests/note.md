# Building Tests

If you want to refactor, the essential precondition is having solid tests

Even if you are fortunate enough to have a tool that can automate the refactorings, you still need tests.

## The Value of Self-Testing Code

+ Those programmers can tell a story of a bug that took a whole day (or more) to find. Fixing the bug is usually pretty quick, but finding it is a nightmare.

+ I’ve found life to be much easier if I have a test suite to lean on.

+ As the Red/Green/Refactor movement advocates, one of the most useful times to write tests is before you start programming. When you need to add a feature, begin by writing the test.

+ **Writing the test also concentrates on the interface rather than the implementation (which is always a good thing).**

+ This book is about refactoring. Refactoring requires tests. If you want to refactor, you have to write tests.

+ But often when I’m working with people on refactoring, we have a body of non-self-testing code to work on. **So first we have to make the code self-testing before we refactor.**



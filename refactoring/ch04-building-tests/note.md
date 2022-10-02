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

+ Again I’ll mention that when I’m writing tests, **I start by making them fail**. With existing code I either change it to make it fail (if I can touch the code) or put an incorrect expected value in the assertion.
    + I do this because I like to prove to myself that the test does actually run and the test is actually testing what it’s supposed to (which is why I prefer changing the tested code if I can).
    + This may be paranoia, but you can really confuse yourself when tests are testing something other than what you think they are testing.

## Developer and Quality Assurance Tests

+ I write them to improve my productivity as a pro- grammer. Making the quality assurance department happy is just a side effect.

+ Quality assurance tests are a different animal. They are written to ensure the software as a whole works. They provide quality assurance to the customer and don’t care about programmer productivity.

+ Tip: **When you get a bug report, start by writing a unit test that exposes the bug.**

## Adding More Tests

+ The key is to test the areas that you are most worried about going wrong. That way you get the most benefit for your testing effort.

+ Tip: Think of the boundary conditions under which things might go wrong and concentrate your tests there.
    + Part of looking for boundaries is looking for special conditions that can cause the test to fail.

+ In the process you get to really understand the interface of the class.
+ In particular, it helps to think about error conditions and boundary conditions.
+ That’s another advantage for writing tests as you write code, or even before you write the production code.

+ **You should concentrate on where the risk is**. Look at the code and see where it becomes complex. Look at the function and consider the likely areas of error. Your tests will not find every bug, but **as you refactor you will understand the program better and thus find more bugs**.

+ Build a good bug detector and run it frequently.




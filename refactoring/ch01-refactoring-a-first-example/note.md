## Removing Temps

+ As I suggested before, temporaty variables can be a problem. They are useful only within their own routine, an thus they encourage long, complex routines.
    + solution: Replace Temps with `Query` / `Query Method`

+ As with `Replace Temp with Query`, this change can cause performance worries to inexperienced programmers. The same advice applies: make the code clean first and then use a profiler to deal with performance issues

+ "I'm now at the point where I take of my refactoring hat and put on my adding function hat"


## Replacing the Conditional Logic with Polymorphism

+ If you must use a case statement, it should be on your own data, not on someone else's

+ The refactoring I'm going to use here is `Replace Type Code with State/Stragety`

## Final thoughts

+ The most important lesson from this example is the rhythm of refactoring: test, small changes, test, small changes, test, small changes. It's that rhythm that allows refactoring to move quickly and safely

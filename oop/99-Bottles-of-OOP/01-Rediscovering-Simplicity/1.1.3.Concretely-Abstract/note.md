1. How difficult was it to write?

难

2. How hard is it to understand?

难

3. How expensive will it be to change?

单独修改都容易, 但修改一个方法会影响到很多其他的方法

+ This solution is characterized by having many small methods.
+ This is normally a good thing, but somehow in this case it’s gone badly wrong.

这种解法使用了很多小的方法, 通常这种做法是有益的, 但这里确实错的


+ It’s obvious that the author of this code was committed to doing the right thing, and that they carefully followed the Red, Green, Refactor style of writing code.
+ The various strings that make up the song are never repeated —it looks as though these strings were refactored into separate methods at the first sign of duplication.

+ The code is DRY, and DRYing out code should save you money.
+ DRY promises that if you put a chunk of code into a method and then invoke that method instead of duplicating the code, you will save money later if the behavior of that chunk changes.

这些代码是dry的, 没有冗余, 减少代码的冗余能为你省钱

+ When you invoke a method instead of implementing behavior, you add a level of indirection.
+ This indirection makes the details of what’s happening harder to understand, but DRY promises that in return, your code will be easier to change.

当用调用一个方法(library)替代自己实现时, 实际上增加了一层间接性

这个间接性会让代码变得更难于理解

但是DRY带来的好处是 它使得代码更易于修改

+ TODO page 29


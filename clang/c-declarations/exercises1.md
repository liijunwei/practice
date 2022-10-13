https://cseweb.ucsd.edu/~ricko/rt_lt.rule.html


```
First, symbols.  Read
     *    as "pointer to"           - always on the left side
     []   as "array of"             - always on the right side
     ()   as "function returning"   - always on the right side
```

STEP 1
------
Find the identifier.  This is your starting point.  Then say to yourself,
"identifier is."  You've started your declaration.


STEP 2
------
Look at the symbols on the right of the identifier.  If, say, you find "()"
there, then you know that this is the declaration for a function.  So you
would then have "identifier is function returning".  Or if you found a
"[]" there, you would say "identifier is array of".  Continue right until
you run out of symbols *OR* hit a *right* parenthesis ")".  (If you hit a
left parenthesis, that's the beginning of a () symbol, even if there
is stuff in between the parentheses.  More on that below.)


STEP 3
------
Look at the symbols to the left of the identifier.  If it is not one of our
symbols above (say, something like "int"), just say it.  Otherwise, translate
it into English using that table above.  Keep going left until you run out of
symbols *OR* hit a *left* parenthesis `(`.


```c
int *p[];
// p is array of poniter to int
// 指针(的)数组, 每个指针指向int类型的数

int (*p)[];
// p is pointer to array of int
// 数组(的)指针, 数组里存放int型的数

```


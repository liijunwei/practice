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

Now repeat steps 2 and 3 until you've formed your declaration.

```c
int *p[];
// p is array of poniter to int
// 指针(的)数组, 每个指针指向int类型的数

int (*p)[];
// p is pointer to array of int
// 数组(的)指针, 数组里存放int型的数

int *(*func())();
// func is function returning pointer to function returning pointer to int

int (*(*fun_one)(char *,double))[9][20];
// fun_one is pointer to function expecting(char *,double) and returning pointer to array(size 9) of array(size 20) of int

int (*(*fun_one)())[][];
// it's not as complicated if you get rid of the array sizes and argument lists
// fun_one is a pointer to function returning pointer to array of array of int

int **p
// p is pointer to pointer to int

int (*p)()
// p is pointer to function returning int

int *p[]
// p is array of pointer to int

int aa[][]
// aa is array of array of int

int af[]() // ILLEGAL, cannot have an array of functions
// af is array of function returning int

int *fp()
// fp is function returning pointer to int

int fa()[] // ILLEGAL, cannot have a function that returns an array
// fa is function returning array of int

int ff()() // ILLEGAL, cannot have a function that returns a function
// ff iss function returning fucntion returning int
```

Some final words:
-------------
It is quite possible to make illegal declarations using this rule,
so some knowledge of what's legal in C is necessary.  For instance,
if the above had been:

```c
int *((*fun_one)())[][];

// it would have been "fun_one is pointer to function returning array of array of
//                                           ^^^^^^^^^^^^^^^^^^^^^^^^
// pointer to int".  Since a function cannot return an array, but only a
// pointer to an array, that declaration is illegal.
```

Illegal combinations include:
```
   []() - cannot have an array of functions
   ()() - cannot have a function that returns a function
   ()[] - cannot have a function that returns an array
```

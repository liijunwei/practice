https://cseweb.ucsd.edu/~ricko/rt_lt.rule.html


```
First, symbols.  Read
     *    as "pointer to"           - always on the left side
     []   as "array of"             - always on the right side
     ()   as "function returning"   - always on the right side
```

```c
int *p[];
// p is an array of pointer to int
// 指针(的)数组, 每个指针指向int类型的数

int (*p)[];
// p is a pointer to an array of ints
// 数组(的)指针, 数组里存放int型的数

```


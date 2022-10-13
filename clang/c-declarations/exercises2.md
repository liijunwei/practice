https://bbs.huaweicloud.com/blogs/307999

```c
int (*p1)(int*, int (*f)(int*));
// p1 is pointer to function expecting(pointer to int and pointer to function returning int) and returning int

int (*p2[5])(int*);
// p2 is array of pointer to function returning int

int (*(*p3)[5])(int*);
// p3 is pointer to array of pointer to function expecting pointer to int and returning int

int*(*(*p4)(int*))(int*);
// p4 is pointer to function returning pointer to function returning pointer to int
// wtf

int (*(*p5)(int*))[5];
// p5 is pointer to function returning pointer to array of int
```

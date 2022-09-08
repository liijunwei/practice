https://bbs.huaweicloud.com/blogs/307999

```c
int (*p1)(int*, int (*f)(int*));

int (*p2[5])(int*);

int (*(*p3)[5])(int*);

int*(*(*p4)(int*))(int*);

int (*(*p5)(int*))[5];
```

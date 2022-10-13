https://www.bilibili.com/video/BV1mB4y1L7HB?share_source=copy_web

```c
int a[]
// a is array of int

int *v[]
// v is array of pointer to int

int (*v)[]
// v is pointer to array of int

int func()
// func is function returning int

int (*func)()
// func is pointer to function returning int

int (*v[])()
// v is array of pointer to function returning int

int (*(*v)[])()
// v is pointer to array of pointer to function returning int
```

extra

```c
int const a
// 等价于
const int a

int const *r;
const int *r;
// r is pointer to const int
// r can be changed
// (*r) can't be changed

int *const r;
// r is const pointer to int
// r can't be changed
// (*r) can be changed
```

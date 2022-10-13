+ OK ./exercises4.md
+ OK ./exercises1.md
+ TODO ./exercises2.md
+ TODO ./exercises3.md

## why

c语言里的 变量 和 函数声明真的太难懂啦


[初中生也能看懂的C/C++类型声明规则教学, 很简单的!](https://www.bilibili.com/video/BV1mB4y1L7HB?share_source=copy_web) -> 右左法则 -> [The "right-left" rule is a completely regular rule for deciphering C declarations.  It can also be useful in creating them.](https://cseweb.ucsd.edu/~ricko/rt_lt.rule.html) -> [Right Left Rule Explanation](https://youtu.be/3zwNIPlrZHQ)

+ https://bbs.huaweicloud.com/blogs/307999

## example.1 `char **argv` <=> `char *argv[]`

+ [What does char * argv[] means?](https://stackoverflow.com/questions/16666353/what-does-char-argv-means#:~:text=The%20declaration%20char%20*argv%5B%5D,a%20pointer%20as%20an%20array)

> The declaration `char *argv[]` is an array (of undetermined size) of pointers to char, in other words an array of strings.
> And all arrays decays(衰变) to pointers, and so you can use an array as a pointer (just like you can use a pointer as an array).
> So `*++argv` first increases the "pointer" to point to the next entry in the array argv (which the first time in the loop will be the first command line argument) and dereferences that pointer.


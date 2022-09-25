#include <stdio.h>

/*
page 86

指针和整数之间不能相互转换, 但0是唯一的例外: 常量0可以赋值给指针,
指针也可以和常量0进行比较 程序中常用符号常量NULL代替0,
这样便于更清晰地说明常量0是指针的一个特殊值

符号常量NULL定义在标准头文件 <stddef.h> 中
*/

#define ALLOCSIZE 10000 // 可用空间大小

static char allocbuf[ALLOCSIZE]; // alloc使用的存储区
static char *allocp = allocbuf;  // 下一空闲位置

// 返回指向n个字符的指针
char *alloc(int n) {
  if (allocbuf + ALLOCSIZE - allocp >= n) { // 有足够的空闲空间
    allocp += n;
    printf("allocating memory: %d\n", n);
    return allocp -
           n; // 明白了: allocp 这个指针在向右移动了n个单位后, 到了新的位置,
              // 返回一个新的位置(指针), 表示分配好的这n个单位内存空间
  } else {
    printf("memory not enough(%d > %d)\n", n, ALLOCSIZE);
    return NULL; // c语言保证, 0永远不是有效的数据地址,
                 // 因此可以用做没有足够空间的标识
  }
}

// 释放p指向的存储区
void afree(char *p) {
  if (p >= allocbuf && p < allocbuf + ALLOCSIZE) {
    printf("freeing up memory\n");
    allocp = p;
  }
}

int main(int argc, char const *argv[]) {

  char *demo1 = alloc(100);
  afree(demo1);

  char *demo2 = alloc(ALLOCSIZE + 1);
  afree(demo2);

  return 0;
}

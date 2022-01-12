#include <stdio.h>

/*
page 86


*/

#define ALLOCSIZE 10000 // 可用空间大小

static char allocbuf[ALLOCSIZE]; // alloc使用的存储区
static char *allocp = allocbuf;  // 下一空闲位置

// 返回指向n个字符的指针
char *alloc(int n){
  if(allocbuf + ALLOCSIZE - allocp >= n){ // 有足够的空闲空间
    allocp += n;
    printf("allocating memory: %d\n", n);
    return allocp - n; // TODO 这一行没懂
  } else {
    printf("memory not enough\n");
    return 0;
  }
}

// 释放p指向的存储区
void afree(char *p){
  if(p >= allocbuf && p < allocbuf + ALLOCSIZE){
    printf("freeing up memory\n");
    allocp = p;
  }
}

int main(int argc, char const *argv[])
{

  char *demo1 = alloc(100);
  afree(demo1);

  char *demo2 = alloc(ALLOCSIZE + 1);
  afree(demo2);

  return 0;
}

#include <stdio.h>

/*
page 86


*/

#define ALLOCSIZE 10000 // 可用空间大小

static char allocbuf[ALLOCSIZE]; // alloc使用的存储区
static char *allocp = allocbuf;  // 下一空闲位置

char *alloc(int n){
  if(allocbuf + ALLOCSIZE - allocp >= n){ // 有足够的空闲空间
    allocp += n;
    return allocp - n;
  } else {
    return 0;
  }
}

void afree(char *p){
  if(p >= allocbuf && p < allocbuf + ALLOCSIZE){
    allocp = p;
  }
}

int main(int argc, char const *argv[])
{

  char *demo = alloc(100);
  afree(demo);

  return 0;
}

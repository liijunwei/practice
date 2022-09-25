#include <stdio.h>

/*
page 89

指针的算数运算具有一致性:
如果处理的数据累心格式比字符型占用更多存储空间的浮点类型,
并且p是一个指向浮点类型的指针 那么在执行p++后, p将指向下一个浮点数的地址;

因此只需要将alloc和afree函数里的char类型替换为float类型,
就可以得到一个适用于浮点类型而非字符型的内存分配函数

所有的指针运算都会自动考虑他所指向的对象的长度
*/

#define ALLOCSIZE 10000

static float allocbuf[ALLOCSIZE];
static float *allocp = allocbuf;

float *alloc(int n) {
  if (allocbuf + ALLOCSIZE - allocp >= n) {
    allocp += n;
    printf("allocating memory: %d\n", n);
    return allocp - n;
  } else {
    printf("memory not enough(%d > %d)\n", n, ALLOCSIZE);
    return NULL;
  }
}

void afree(float *p) {
  if (p >= allocbuf && p < allocbuf + ALLOCSIZE) {
    printf("freeing up memory\n");
    allocp = p;
  }
}

int main(int argc, char const *argv[]) {

  float *demo1 = alloc(100);
  afree(demo1);

  float *demo2 = alloc(ALLOCSIZE + 1);
  afree(demo2);

  return 0;
}

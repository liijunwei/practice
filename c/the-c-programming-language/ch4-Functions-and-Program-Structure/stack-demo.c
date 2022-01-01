#include <stdio.h>
#include <assert.h>

/*
page 65

stack data structure
*/

#define MAXVAL 5  // 栈val的最大深度
int sp = 0;         // 下一个空闲栈的位置
double val[MAXVAL]; // 值栈

// 把f压入栈中
void push(double f){
  if(sp < MAXVAL){
    val[sp] = f;
    sp++;
  } else {
    printf("Error: stack(size: %d) full, can't push %g\n", MAXVAL, f);
  }
}

// 弹出并返回栈顶的值
double pop(void){
  if(sp > 0) {
    return val[--sp];
  } else {
    printf("Error: stack empty\n");
    return 0.0;
  }
}

int main(int argc, char const *argv[])
{
  push(1);
  push(2);
  push(3);
  push(4);
  push(5);
  push(6);

  for(int i = sp - 1; i >= 0; i--){
    printf("%f\n", val[i]);
  }

  return 0;
}

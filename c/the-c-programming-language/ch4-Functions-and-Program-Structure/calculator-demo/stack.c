#include <stdio.h>
#include "./calc.h"

#define MAXVAL 100 // 栈的最大深度

static double buf[MAXVAL]; // 值栈
static int bufp = 0;       // 下一个空闲栈的位置

// 把f压入栈中
void push(double f){
  if(bufp < MAXVAL){
    buf[bufp] = f;
    bufp++;
  } else {
    printf("Error: stack full, can't push %g\n", f);
  }
}

// 弹出并返回栈顶的值
double pop(void){
  if(bufp > 0) {
    return buf[--bufp];
  } else {
    printf("Error: stack empty\n");
    return 0.0;
  }
}

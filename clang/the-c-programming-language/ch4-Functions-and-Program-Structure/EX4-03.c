#include <stdio.h>
#include <stdlib.h>
#include <ctype.h>
#include <math.h> /* for fomod */

#include "../common-utils/getch-ungetch.c"

/*
page 67

在有了基本框架后, 对计算器程序进行扩充就比较简单了
在该程序中加入取模(%)运算符, 并注意考虑负数的情况

ch4-Functions-and-Program-Structure/calculator-Reverse-Polish-notation-demo.c

*/


#define MAXVAL 100 // 栈val的最大深度
int sp = 0; // 下一个空闲栈的位置
double val[MAXVAL]; // 值栈

// 把f压入栈中
void push(double f){
  if(sp < MAXVAL){
    val[sp] = f;
    sp++;
  } else {
    printf("Error: stack full, can't push %g\n", f);
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

#define MAXOP  100

int main(int argc, char const *argv[])
{
  int type;
  double op2;
  char s[MAXOP];

  while((type = getop(s)) != EOF){
    switch(type){
      case NUMBER:
        push(atof(s));
        break;
      case '+':
        push(pop() + pop());
        break;
      case '*':
        push(pop() * pop());
        break;
      case '-':
        op2 = pop();
        push(pop() - op2);
        break;
      case '/':
        op2 = pop();
        if(op2 != 0.0){
          push(pop() / op2);
        } else {
          printf("Error: zero devisor\n");
        }
        break;
      case '%':
        op2 = pop();
        if(op2 != 0.0){
          push(fmod(pop(), op2));
        } else {
          printf("Error: zero devisor\n");
        }
        break;
      case '\n':
        printf("Result:\t%.8g\n", pop());
        break;
      default:
        printf("Error: unknown command %s\n", s);
        break;
    }
  }

  return 0;
}

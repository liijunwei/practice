/*
page 140

改写第四章的后缀计算器, 用scanf函数和(或)sscanf函数实现输入以及数的转换

ref:
  ch4-Functions-and-Program-Structure/EX4-03.c

思路:
  只需要修改getop函数, 在调用getop时, 字符总是跟在一个数字后面出现;
  lastc是具有两个元素的静态数组, 用于记录最后读入的字符(sscanf读入的一个字符串)
*/

#include <stdio.h>
#include <stdlib.h>
#include <ctype.h>
#include <math.h> /* for fomod */

#define NUMBER  '0' // signal that a number was found
#define MAXVAL 100  // 栈val的最大深度
int sp = 0;         // 下一个空闲栈的位置
double val[MAXVAL]; // 值栈

/* get next operator or numeric operend */
int getop(char s[]) {
  char c;
  int i;
  int rc;
  static char lastc[] = " ";

  sscanf(lastc, "%c", &c);
  lastc[0] = ' ';

  while ((s[0] = c) == ' ' || c == '\t') {
    if (scanf("%c", &c) == EOF) {
      c = EOF;
    }
  }

  s[1] = '\0';

  if (!isdigit(c) && c != '.') {
    return c;
  }

  i = 0;

  if (isdigit(c)) {
    do {
      rc = scanf("%c", &c);
      if (!isdigit(s[++i] = c)) {
        break;
      }
    } while (rc != EOF);
  }

  if (c == '.') {
    do {
      rc = scanf("%c", &c);
      if (!isdigit(s[++i] = c)) {
        break;
      }
    } while (rc != EOF);
  }

  s[i] = '\0';

  if (rc != EOF) {
    lastc[0] = c;
  }

  return NUMBER;
}

void push(double f){
  if(sp < MAXVAL){
    val[sp] = f;
    sp++;
  } else {
    printf("Error: stack full, can't push %g\n", f);
  }
}

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

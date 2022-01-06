#include <stdio.h>
#include <stdlib.h>
#include <ctype.h>

/*
page 71

修改getop函数, 使其不必使用ungetch函数
提示: 可以使用一个static类型的内部变量解决该问题

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

#define BUFSIZE 100
char buf[BUFSIZE];
int bufp = 0;

int getch(void){
  return (bufp > 0) ? buf[--bufp] : getchar();
}

#define NUMBER '0'
int getop(char s[]){
  int i;
  int c;
  static int lastc = 0; // here

  if(lastc == 0){
    c = getch();
  } else {
    c = lastc;
    lastc = 0;
  }

  while((s[0] = c) == ' ' || c == '\t'){
    c = getch(); // here
  }

  s[1] = '\0';
  if(!isdigit(c) && c != '.') {
    return c; // 不是数
  }

  i = 0;
  if(isdigit(c)) {
    // 整数部分
    while(isdigit(s[++i] = c = getch())){
      ;
    }
  }

  if(c == '.') {
    // 小数部分
    while(isdigit(s[++i] = c = getch())){
      ;
    }
  }

  s[i] = '\0';

  if(c != EOF) {
    lastc = c; // here
  }

  return NUMBER;
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

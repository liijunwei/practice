#include <ctype.h>
#include <stdio.h>
#include <stdlib.h>

/*
page 63

编写一个具有 加减乘除 四则运算功能的计算程序
使用逆波兰表达式代替普通的中缀表示法
*/

#define MAXVAL 100  // 栈val的最大深度
int sp = 0;         // 下一个空闲栈的位置
double val[MAXVAL]; // 值栈

// 把f压入栈中
void push(double f) {
  if (sp < MAXVAL) {
    val[sp] = f;
    sp++;
  } else {
    printf("Error: stack full, can't push %g\n", f);
  }
}

// 弹出并返回栈顶的值
double pop(void) {
  if (sp > 0) {
    return val[--sp];
  } else {
    printf("Error: stack empty\n");
    return 0.0;
  }
}

#define BUFSIZE 100
char buf[BUFSIZE];
int bufp = 0; // buf中下一空闲的位置

// 取一个字符(可能是压回的字符)
int getch(void) { return (bufp > 0) ? buf[--bufp] : getchar(); }

// 把字符压回输入中
void ungetch(int c) {
  if (bufp >= BUFSIZE) {
    printf("Ungetch, too many characters\n");
  } else {
    buf[bufp++] = c;
  }
}

#define NUMBER '0'

int getop(char s[]) {
  int i;
  int c;

  while ((s[0] = c = getch()) == ' ' || c == '\t') {
    ;
  }

  s[1] = '\0';
  if (!isdigit(c) && c != '.') {
    return c; // 不是数
  }

  i = 0;
  if (isdigit(c)) {
    // 整数部分
    while (isdigit(s[++i] = c = getch())) {
      ;
    }
  }

  if (c == '.') {
    // 小数部分
    while (isdigit(s[++i] = c = getch())) {
      ;
    }
  }

  s[i] = '\0';

  if (c != EOF) {
    ungetch(c);
  }

  return NUMBER;
}

#define MAXOP 100

int main(int argc, char const *argv[]) {
  int type;
  double op2;
  char s[MAXOP];

  while ((type = getop(s)) != EOF) {
    switch (type) {
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
      if (op2 != 0.0) {
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

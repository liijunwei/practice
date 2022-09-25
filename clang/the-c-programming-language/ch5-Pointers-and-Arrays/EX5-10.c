#include <stdio.h>
#include <stdlib.h> // atof

#include "../common-utils/getch-ungetch.c" // getch/ungetch/getop/ungets
#include "../common-utils/stack.c"         // push/pop

/*
page 102

编写程序expr, 计算命令行输入的米波兰表达式的值,
其中每个运算符或操作数用一个单独的参数表示 例如: `expr 2 3 4 + *` 将计算 2 x (3
+ 4) 的值

*/

#define MAXOP 100 // max size of operand or operator

/*
./a.out 2 3 4 + +
Result: 9

* 号在命令行里得转义(吗)
./a.out 2 3 4 + \*
Result: 14
*/

// reverse polish calculator; uses command line
int main(int argc, char *argv[]) {
  char s[MAXOP];
  double op2;

  while (--argc > 0) {
    ungets(" ");
    ungets(*++argv);
    switch (getop(s)) {
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

  printf("Result: %.8g\n", pop());

  return 0;
}

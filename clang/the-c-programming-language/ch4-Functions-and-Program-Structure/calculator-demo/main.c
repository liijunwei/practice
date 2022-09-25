#include <stdio.h>
#include <stdlib.h>

#include "./calc.h"

#define MAXOP 100

// page 70
// make

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

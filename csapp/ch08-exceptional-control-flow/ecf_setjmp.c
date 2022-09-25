#include "csapp.h"

jmp_buf buf;

int error1 = 0;
int error2 = 1;

void foo();
void bar();

int main(int argc, char const *argv[]) {
  switch (setjmp(buf)) { // 什么意思??
  case 0:
    foo();
    break;
  case 1:
    printf("Detected an error1 condition in foo\n");
    break;
  case 2:
    printf("Detected an error2 condition in foo\n");
    break;
  default:
    printf("Unknown error condition in foo\n");
  }

  exit(0);
}

// deeply nested function foo
void foo() {
  if (error1) {
    longjmp(buf, 1);
  }

  bar();
}

void bar() {
  if (error2) {
    longjmp(buf, 2);
  }
}

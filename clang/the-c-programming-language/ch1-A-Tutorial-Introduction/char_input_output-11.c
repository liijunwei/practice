#include <stdio.h>

#define CONST_CHAR_SPACE ' '

// 将输入复制到输出, 并将其中连续的多个空格用一个空格代替

// 变量.1 char          是/不是 ' '
// 变量.2 space_counter 是/不是 0
// 两两组合, 四种情况

int main() {
  int c;
  printf("please enter chars, `ctrl+d` to exit\n");

  int space_counter = 0;

  while ((c = getchar()) != EOF) {
    if (c == CONST_CHAR_SPACE && space_counter == 0) {
      putchar(c);
      space_counter += 1;
    }

    if (c != CONST_CHAR_SPACE) {
      putchar(c);
      space_counter = 0;
    }
  }
}

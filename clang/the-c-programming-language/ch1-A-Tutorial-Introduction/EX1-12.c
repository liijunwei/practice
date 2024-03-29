#include <stdio.h>

#define CONST_IN 1  // 在单词内
#define CONST_OUT 0 // 在单词外

#define CONST_CHAR_TAB '\t'
#define CONST_CHAR_NEWLINE '\n'
#define CONST_CHAR_SPACE ' '

// page 15 练习1-12

// 情景1 是否为空格
// 情景2 是否在单词内
// 两两组合, 4种情况

// 编写程序, 以每行一个单词的形式打印其输入

int main() {
  int c;
  int state = CONST_OUT;

  printf("please enter chars, `ctrl+d` to exit\n");

  while ((c = getchar()) != EOF) {
    if (c == CONST_CHAR_SPACE || c == CONST_CHAR_NEWLINE ||
        c == CONST_CHAR_TAB) {
      if (state == CONST_IN) { // ending of a word
        putchar(CONST_CHAR_NEWLINE);
        state = CONST_OUT;
      }
    } else if (state == CONST_OUT) { // beginning of a word
      state = CONST_IN;
      putchar(c);
    } else { // inside a word
      putchar(c);
    }
  }
}

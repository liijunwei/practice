#include <stdio.h>

// page 17
// 编写程序, 打印输入中单词长度的直方图. 水平方向的直方图比较容易绘制,
// 垂直方向的直方图则要困难一些

// 示例:
// 单词长度为 1 的有 3 个
// 单词长度为 2 的有 3 个
// 单词长度为 3 的有 1 个
// 单词长度为 4 的有 10 个
// 单词长度为 5 的有 5 个
// 单词长度为 6 的有 8 个

// 参考: ch1-A-Tutorial-Introduction/array-01.c
// 参考: ch1-A-Tutorial-Introduction/word_count-01.c

// 粗糙得完成了
//    没处理overflow
//    没处理垂直打印
//    ...

#define MAX_WORD_LENGTH 15
#define CONST_IN 1  // 在单词内
#define CONST_OUT 0 // 在单词外

int main() {

  int i;
  int c;
  int word_length[MAX_WORD_LENGTH];
  int num_chars = 0;
  int cursor_state = CONST_OUT;

  printf("Array Initializing...\n");
  for (i = 0; i < MAX_WORD_LENGTH; ++i) {
    word_length[i] = 0;
  }
  printf("Array Initialized!\n");
  printf("please enter chars, `ctrl+d` to exit\n");

  while ((c = getchar()) != EOF) {
    if (c == ' ' || c == '\n' || c == '\t') {
      cursor_state = CONST_OUT;
      if (num_chars > 0) {
        ++word_length[num_chars];
      }

      num_chars = 0;
    } else if (cursor_state == CONST_OUT) {
      cursor_state = CONST_IN;
      num_chars = 1;
    } else {
      ++num_chars;
    }
  }

  for (int i = 0; i < MAX_WORD_LENGTH; i++) {
    printf("%d: ", i);
    for (int j = 0; j < word_length[i]; j++) {
      putchar('*');
    }
    printf("\n");
  }
}

#include <stdio.h>

// 编写程序, 打印输入中单词长度的直方图. 水平方向的直方图比较容易绘制, 垂直方向的直方图则要困难一些

// 示例:
// 单词长度为 1 的有 3 个
// 单词长度为 2 的有 3 个
// 单词长度为 3 的有 1 个
// 单词长度为 4 的有 10 个
// 单词长度为 5 的有 5 个
// 单词长度为 6 的有 8 个

#define MAX_WORD_LENGTH 15
#define ARRAY_SIZE 15

int main(){

  int i;
  int c;
  int word_length[MAX_WORD_LENGTH];

  printf("Array Initializing...\n");
  for(i = 0; i < MAX_WORD_LENGTH; ++i){
    word_length[i] = 0;
  }
  printf("Array Initialized!\n");
  printf("please enter chars, `ctrl+d` to exit\n");

  while((c = getchar()) != EOF){
    putchar(c);
  }
}

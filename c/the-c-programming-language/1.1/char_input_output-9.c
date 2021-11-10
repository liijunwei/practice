#include <stdio.h>

#define CONST_CHAR_SPACE   ' '
#define CONST_CHAR_TAB     '\t'
#define CONST_CHAR_NEWLINE '\n'

// 将输入复制到输出, 并将其中连续的多个空格用一个空格代替

// 步骤.1 识别出连续的多个空格
// 步骤.2 遇到连续的空格时, 输出空, 或者不输出
int main(){
  int c;
  printf("please enter chars, `ctrl+d` to exit\n");

  while((c = getchar()) != EOF){
    putchar(c);
  }
}

#include <stdio.h>

#define CONST_CHAR_SPACE   ' '
#define CONST_CHAR_TAB     '\t'
#define CONST_CHAR_NEWLINE '\n'

// 将输入复制到输出, 并将其中连续的多个空格用一个空格代替
int main(){
  int c;
  printf("please enter chars, `ctrl+d` to exit\n");

  while((c = getchar()) != EOF){
    putchar(c);
  }
}

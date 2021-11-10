#include <stdio.h>

#define CONST_CHAR_TAB       '\t'
#define CONST_CHAR_NEWLINE   '\n'
#define CONST_CHAR_BACKSPACE '\b'
#define CONST_CHAR_SLASH     '\\'

// 将输入复制到输出, 并将其中 的制表符替换为 \t, 回退符替换为 \b, 反斜杠替换为 \\
// 这样就可以将制表符和回退符以可见的形式显示出来

int main(){
  int c;
  printf("please enter chars, `ctrl+d` to exit\n");

  while((c = getchar()) != EOF){
    if(c == CONST_CHAR_TAB){
      printf("\\t");
    }

    if(c == CONST_CHAR_NEWLINE){
      printf("\\n");
    }

    if(c == CONST_CHAR_BACKSPACE){
      printf("\\b");
    }

    if(c == CONST_CHAR_SLASH){
      printf("\\\\");
    }

    if(c != CONST_CHAR_TAB       &&
       c != CONST_CHAR_BACKSPACE &&
       c != CONST_CHAR_SLASH
       ){
      putchar(c);
    }
  }
}

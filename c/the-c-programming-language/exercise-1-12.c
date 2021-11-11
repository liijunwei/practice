#include <stdio.h>

#define CONST_IN  1 // 在单词内
#define CONST_OUT 0 // 在单词外

#define CONST_CHAR_TAB       '\t'
#define CONST_CHAR_NEWLINE   '\n'
#define CONST_CHAR_SPACE     ' '


// 编写程序, 以每行一个单词的形式打印其输入

int main(){
  int c;
  int state = CONST_OUT;

  printf("please enter chars, `ctrl+d` to exit\n");

  while((c = getchar()) != EOF){
    if(c == CONST_CHAR_SPACE || c == CONST_CHAR_NEWLINE || c == CONST_CHAR_TAB){
      if(state == CONST_IN){
        putchar(CONST_CHAR_NEWLINE);
        state = CONST_OUT;
      }
    }
    else if(state == CONST_OUT){
      state = CONST_IN;
      putchar(c);
    }
    else{
      putchar(c);
    }
  }

}

#include <stdio.h>

#define CONST_CHAR_SPACE   ' '
#define CONST_CHAR_TAB     '\t'
#define CONST_CHAR_NEWLINE '\n'

int main(){
  int c;
  printf("please enter chars, `ctrl+d` to exit\n");

  while((c = getchar()) != EOF){
    putchar(c);
  }
}

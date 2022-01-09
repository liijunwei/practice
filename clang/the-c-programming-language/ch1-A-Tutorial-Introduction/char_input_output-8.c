#include <stdio.h>

// page 13 练习1-8

#define CONST_CHAR_SPACE   ' '
#define CONST_CHAR_TAB     '\t'
#define CONST_CHAR_NEWLINE '\n'

int main(){
  int c;

  int char_space_count = 0;
  int char_tab_count = 0;
  int char_newline_count = 0;

  printf("please enter chars, `ctrl+d` to exit\n");

  while((c = getchar()) != EOF){
    if(c == CONST_CHAR_SPACE){
      ++char_space_count;
    }

    if(c == CONST_CHAR_TAB){
      ++char_tab_count;
    }

    if(c == CONST_CHAR_NEWLINE){
      ++char_newline_count;
    }
  }

  printf("\n");
  printf("spaces:   %d\n", char_space_count);
  printf("tabs:     %d\n", char_tab_count);
  printf("newlines: %d\n", char_newline_count);
}

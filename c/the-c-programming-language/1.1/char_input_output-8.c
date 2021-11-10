#include <stdio.h>

#define CONST_CHAR_SPACE    ' '
#define CONST_CHAR_TAB      '\t'
#define CONST_CHAR_NEW_LINE '\n'

int main(){
  int c;
  int n1 = 0;

  int char_space_count = 0;
  int char_tab_count = 0;
  int char_new_line_count = 0;

  printf("please enter chars, `ctrl+d` to exit\n");

  while((c = getchar()) != EOF){
    if(c == CONST_CHAR_SPACE){
      ++char_space_count;
    }

    if(c == CONST_CHAR_TAB){
      ++char_tab_count;
    }

    if(c == CONST_CHAR_NEW_LINE){
      ++char_new_line_count;
    }
  }

  printf("\n");
  printf("space chars count is: %d\n", char_space_count);
  printf("tab chars count is: %d\n", char_tab_count);
  printf("new_line chars count is: %d\n", char_new_line_count);
}

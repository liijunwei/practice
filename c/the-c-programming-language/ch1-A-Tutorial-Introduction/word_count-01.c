#include <stdio.h>

// page 21
#define CONST_IN  1 // 在单词内
#define CONST_OUT 0 // 在单词外

#define CONST_CHAR_TAB       '\t'
#define CONST_CHAR_NEWLINE   '\n'
#define CONST_CHAR_SPACE     ' '

// 统计输入的 行数/单词数/字符数

int main(){
  int num_line = 0;
  int num_word = 0;
  int num_char = 0;

  int c;
  int state = CONST_OUT;

  printf("please enter chars, `ctrl+d` to exit\n");

  while((c = getchar()) != EOF){
    ++num_char;

    if(c == CONST_CHAR_NEWLINE){
      ++num_line;
    }

    if(c == CONST_CHAR_SPACE || c == CONST_CHAR_NEWLINE || c == CONST_CHAR_TAB){
      state = CONST_OUT;
    }
    else if(state == CONST_OUT){
      state = CONST_IN;
      ++num_word;
    }
    else{
      // do nothing
    }
  }

  printf("num_line: %d num_word: %d num_char: %d \n", num_line, num_word, num_char);
}

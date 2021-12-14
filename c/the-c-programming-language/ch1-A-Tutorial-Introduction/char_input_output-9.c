#include <stdio.h>

#define CONST_CHAR_SPACE ' '

// page 13 练习1-9

// 将输入复制到输出, 并将其中连续的多个空格用一个空格代替

// https://blog.csdn.net/go_to_break/article/details/76944018?spm=1001.2101.3001.6650.1&utm_medium=distribute.pc_relevant.none-task-blog-2%7Edefault%7ECTRLIST%7Edefault-1.no_search_link&depth_1-utm_source=distribute.pc_relevant.none-task-blog-2%7Edefault%7ECTRLIST%7Edefault-1.no_search_link

// 变量.1 char          是/不是 ' '
// 变量.2 space_counter 是/不是 0
// 两两组合, 四种情况

int main(){
  int c;
  printf("please enter chars, `ctrl+d` to exit\n");

  int space_counter = 0;

  while((c = getchar()) != EOF){
    if(c != CONST_CHAR_SPACE){
      space_counter = 0;
      putchar(c);
    }

    if(c == CONST_CHAR_SPACE && space_counter == 0){
      putchar(c);
      space_counter += 1;
    } else {
      // do nothing
    }

  }
}

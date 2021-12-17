#include <stdio.h>
#include <string.h>

// page 25

/*
编写程序entab, 将空格串替换为最少数量的制表符和空格, 但要保持单词之间的间隔不变
假设制表符终止位的位置与练习1-20的detab程序的情况相同。
当使用一个制表符或者一个空格都可以到达下一个制表符终止位时, 选用哪种替换字符比较好
*/

// TODO 没明白

/*
程序的主要想法是: 找出全部空格, 变量pos每递增到
*/

#define TABING 8 // tab increment size
#define CHAR_BLANK 'b' // ' '
#define CHAR_TAB   't' // '\t'

// replace strings of blanks with tabs and blanks
int main(){
  int c;
  int nb  = 0; // number of blanks necessary
  int nt  = 0; // number of tabs necessary
  int pos;     // position of character in line

  for(pos = 1; (c = getchar()) != EOF; ++pos){
    if(c == ' '){
      if(pos % TABING != 0){
        ++nb;
      } else {
        nb = 0;
        ++nt;
      }
    } else {
      for(; nt > 0; --nt){
        putchar(CHAR_TAB);
      }

      if(c == 't'){
        nb = 0;
      } else {
        for(; nb > 0; --nb){
          putchar(CHAR_BLANK);
        }
      }

      putchar(c);

      if(c == '\n'){
        pos = 0;
      } else if(c == '\t'){
        pos = pos + (TABING - (pos - 1) % TABING) - 1;
      }
    }
  }

  return 0;
}

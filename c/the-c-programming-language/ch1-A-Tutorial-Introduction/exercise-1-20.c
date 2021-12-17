#include <stdio.h>
#include <string.h>

// page 25

// 编写程序detab, 将输入中的制表符替换成适当数目的空格, 使空格充满到下一个制表符终止位的地方
// 假设制表符终止位的位置是固定的, 比如每个n列就会出现一个制表符终止位, n应该为变量还是符号常量呢?

#define TABING 8 // tab increment size

// replace tabs with the proper number of blanks
int main(){
  int c;
  int nb  = 0; // number of blanks necessary
  int pos = 1; // position of character in line

  while((c = getchar()) != EOF){
    if(c == '\t'){
      nb = TABING - (pos - 1) % TABING;
      while(nb > 0){
        putchar(' ');
        ++pos;
        --nb;
      }
    } else if(c == '\n'){
      putchar(c);
      pos = 1;
    } else {
      putchar(c);
      ++pos;
    }
  }

  return 0;
}

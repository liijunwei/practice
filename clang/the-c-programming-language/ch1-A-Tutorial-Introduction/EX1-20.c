#include <stdio.h>
#include "../common-utils/getline.c"

// page 25

// 懂了 :D
// 编写程序detab, 将输入中的制表符替换成适当数目的空格, 使空格充满到下一个制表符终止位的地方
// 假设制表符终止位的位置是固定的, 比如每个n列就会出现一个制表符终止位, n应该为变量还是符号常量呢?(本例中, 使用常量)

// ch5-Pointers-and-Arrays/EX5-11.c
// OK 如何检查程序的正确性?
// 将空格 ' ' 换为其他可视的符号即可, 比如 '*'

#define TABING 8 // tab increment size
#define VISIABLE_CHAR '*'
#define MAXLINE 1000 // 允许的输入行的最大长度

void detab(char src[], char tar[]);

// replace tabs with the proper number of blanks
int main(){
  char s[MAXLINE];
  char t[MAXLINE * TABING];

  printf("Note: for the sake of clarity, all tabs will be replaced with '%c'.\n", VISIABLE_CHAR);
  while(custom_getline(s, MAXLINE) > 0) {
    detab(s, t);
    printf("%s\n", t);
  }

  return 0;
}

void detab(char src[], char tar[]){
  int c;
  int nb  = 0; // number of blanks necessary

  int i;
  int j = 0;

  for(i = 0; (c = src[i]) != '\0'; i++){
    if(c == '\t'){
      nb = TABING - j % TABING;
      while(nb > 0){
        tar[j] = VISIABLE_CHAR;
        nb--;
        j++;
      }
    } else {
      tar[j] = c;
      j++;
    }
  }

  tar[j] = '\0';
}

#include <stdio.h>

// 将输入复制到输出 V1
int main(){
  int c;

  c = getchar();
  while(c != EOF){
    putchar(c);
    c = getchar();
  }

}
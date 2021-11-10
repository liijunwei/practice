#include <stdio.h>

// 对字符进行计数
int main(){
  long nc = 0;

  while(getchar() != EOF){
    ++nc;
  }

  printf("char count is %ld\n", nc);
}

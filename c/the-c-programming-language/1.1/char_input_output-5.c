#include <stdio.h>

// 对字符进行计数
int main(){
  long nc = 0;

  printf("please enter chars, `ctrl+d` to exit\n");
  printf("\n");

  while(getchar() != EOF){
    ++nc;
  }

  printf("char count is %ld\n", nc);
}

#include <stdio.h>

// 行计数
int main(){
  int c;
  int n1 = 0;

  printf("please enter chars, `ctrl+d` to exit\n");

  while((c = getchar()) != EOF){
    if(c == '\n'){
      ++n1;
    }
  }

  printf("line count is: %d\n", n1);
}

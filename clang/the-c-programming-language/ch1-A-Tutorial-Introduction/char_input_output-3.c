#include <stdio.h>

// page 11
// 验证表达式 getchar() != EOF 的值是 0 还是 1
int main(){
  int c;

  // != 优先级比 = 高
  while(c = (getchar() != EOF)){
    printf("%d\n", c); // 1
  }

  printf("%d\n", c); // 0
}

#include <stdio.h>

// 验证表达式 getchar() != EOF 的值是 0 还是 1
int main(){
  int c;
  c = (getchar() != EOF);
  printf("%d", c); // 1
}

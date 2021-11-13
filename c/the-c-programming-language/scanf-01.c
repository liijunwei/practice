#include <stdio.h>

#define CONST_CHAR_SPACE ' '

// 将输入复制到输出, 并将其中连续的多个空格用一个空格代替

// 变量.1 char          是/不是 ' '
// 变量.2 space_counter 是/不是 0
// 两两组合, 四种情况

int main(){

  int a,b,c;

  printf("请输入三个数字：");
  scanf("%d%d%d",&a,&b,&c);
  printf("%d,%d,%d\n",a,b,c);
  return 0;
}

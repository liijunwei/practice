#include <stdio.h>
#include <string.h>
#include <assert.h>

/*
page 53

编写函数 itob(n, s, b), 将整数n转换为以b为底的数, 并将转换结果以字符的形式保存到字符串s中
例如: itob(n, s, 16) 把整数n格式化为16进制的整数保存到s中

*/

#define abs(x) ((x) < 0 ? -(x) : (x))

// ch3-Control-Flow/reverse-demo.c
void reverse(char s[]){
  int c;
  int i;
  int j;

  for(i = 0, j = strlen(s) - 1; i < j; i++, j--){
    c = s[i];
    s[i] = s[j];
    s[j] = c;
  }
}

// 把数字转为字符串
// 处理不了最小的负数
void itob(int n, char s[], int base){
  int i;
  int j;
  int sign;

  if((sign = n) < 0){
    n = -n;
  }

  i = 0;

  do {
    j = n % base;
    s[i++] = (j <= 9) ? (j + '0') : (j + 'a' - 10);
  } while((n /= base) > 0);

  if(sign < 0){
    s[i++] = '-';
  }

  s[i] = '\0';
  reverse(s);
}

int main(int argc, char const *argv[])
{
  char buffer[100];

  itob(-21474836, buffer, 10);
  assert(strcmp("-21474836", buffer) == 0);

  itob(-21474836, buffer, 16);
  assert(strcmp("-147ae14", buffer) == 0);

  itob(16, buffer, 16);
  assert(strcmp("10", buffer) == 0);

  itob(15, buffer, 16);
  assert(strcmp("f", buffer) == 0);

  printf("PASS.\n");
  return 0;
}

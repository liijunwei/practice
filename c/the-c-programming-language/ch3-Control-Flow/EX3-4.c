#include <stdio.h>
#include <limits.h>
#include <string.h>
#include <assert.h>

/*
page 53

TODO not clear

在数的二进制补码表示中, 我们编写的itoa函数不能处理最大的负数, 即n等于 `-(2^字长 - 1)` 的情况
1. 请解释其原因
  -(2^(字长-1)) 无法通过 n = -n; 转换为一个正数; 原因: 对数的二进制的补码表示中, 最大正数只能是 2^(字长-1) - 1

2. 修改该函数, 使它在任何机器上运行时都能打印出正确的值
  + 思路: 先按逆序求出各位数字, 再用reverse函数对字符串s里的字符做一次颠倒得到最终结果
  + sign保存n的初值, 用宏abs计算n%1的绝对值
  + 把取模操作的结果转为正数, 从而绕过无法把最大负值转为正数的问题

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
void itoa(int n, char s[]){
  int i;
  int sign;

  sign = n;
  i = 0;

  do {
    s[i++] = abs(n % 10) + '0';
  } while((n /= 10) != 0);

  if(sign < 0){
    s[i++] = '-';
  }

  s[i] = '\0';
  reverse(s);
}

int main(int argc, char const *argv[])
{
  // printf("min: %d\n", INT_MIN); // -2147483648
  char buffer[100];

  itoa(-2147483648, buffer);
  assert(strcmp("-2147483648", buffer) == 0);

  printf("PASS.\n");
  return 0;
}

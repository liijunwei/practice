/*
page 137

输入函数scanf对应输入函数printf, 他在与后者相反的方向上提供同样的转换功能
`int scanf(char *fmt, ...)`

另外还有个sscanf函数, 用于从一个字符串(而不是标准输入中读取字符序列)
`int sscanf(char *string, char *fmt, ...)`
*/

#include <stdio.h>

int main(int argc, char const *argv[])
{
  char string1[100];
  char string2[100];
  printf("string1(before): %s\n", string1);
  printf("string2(before): %s\n", string2);

  char sentense[100] = "today is Sunday and a Sunny day, this is the last day(7th) of sping festival holiday";
  sscanf(sentense, "today is %s and a %s day", string1, string2);
  printf("string1(after):  %s\n", string1);
  printf("string2(after):  %s\n", string2);

  int number;
  printf("number(before): %d\n", number);
  sscanf(sentense, "today is Sunday and a Sunny day, this is the last day(%dth) of sping festival holiday", &number);
  printf("number(after):  %d\n", number);

  return 0;
}

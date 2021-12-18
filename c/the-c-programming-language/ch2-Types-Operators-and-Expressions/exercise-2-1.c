#include <stdio.h>
#include <limits.h>

/*
page 28

编写一个程序以确定分别由signed及unsigned限定的char、short、int及long类型变量的取值范围。
采用打印标准头文件中的相应值以及直接计算两种方式实现。
通过直接计算来确定浮点类型的取值范围是一项难度很大的任务
*/

void solu_1();
void solu_2();

int main(int argc, char const *argv[])
{
  solu_1();
  solu_2();

  return 0;
}

void solu_1(){
  printf("Signed Types\n===============\n");
  printf("signed char min  = %d\n", SCHAR_MIN);
  printf("signed char max  = %d\n", SCHAR_MAX);
  printf("signed short min = %d\n", SHRT_MIN);
  printf("signed short max = %d\n", SHRT_MAX);
  printf("signed int min   = %d\n", INT_MIN);
  printf("signed int max   = %d\n", INT_MAX);
  printf("signed long min  = %ld\n", LONG_MIN);
  printf("signed long max  = %ld\n", LONG_MAX);

  printf("\n");

  printf("Unsigned Types\n===============\n");
  printf("unsigned char max  = %u\n", UCHAR_MAX);
  printf("unsigned short max = %u\n", USHRT_MAX);
  printf("unsigned int max   = %u\n", UINT_MAX);
  printf("unsigned long max  = %lu\n", ULONG_MAX);
}

void solu_2(){

}


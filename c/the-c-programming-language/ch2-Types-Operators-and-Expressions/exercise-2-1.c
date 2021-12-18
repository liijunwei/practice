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
  printf("\n\n");

  solu_2();

  return 0;
}

// refer to: https://code.woboq.org/userspace/glibc/include/limits.h.html
void solu_1(){
  printf("Signed Types(%s)\n\n", __func__);
  printf("signed char min  = %d\n", SCHAR_MIN);
  printf("signed char max  = %d\n", SCHAR_MAX);
  printf("signed short min = %d\n", SHRT_MIN);
  printf("signed short max = %d\n", SHRT_MAX);
  printf("signed int min   = %d\n", INT_MIN);
  printf("signed int max   = %d\n", INT_MAX);
  printf("signed long min  = %ld\n", LONG_MIN);
  printf("signed long max  = %ld\n", LONG_MAX);

  printf("\n");

  printf("Unsigned Types(%s)\n\n", __func__);
  printf("unsigned char max  = %u\n", UCHAR_MAX);
  printf("unsigned short max = %u\n", USHRT_MAX);
  printf("unsigned int max   = %u\n", UINT_MAX);
  printf("unsigned long max  = %lu\n", ULONG_MAX);
}

// TODO 没懂...
void solu_2(){
  printf("Signed Types(%s)\n\n", __func__);
  printf("signed char min   = %d\n", -(char)((unsigned char) ~0 >> 1));
  printf("signed char max   = %d\n",  (char)((unsigned char) ~0 >> 1));
  printf("signed short min  = %d\n", -(short)((unsigned short) ~0 >> 1));
  printf("signed short max  = %d\n",  (short)((unsigned short) ~0 >> 1));
  printf("signed int min    = %d\n", -(int)((unsigned int) ~0 >> 1));
  printf("signed int max    = %d\n",  (int)((unsigned int) ~0 >> 1));
  printf("signed long min   = %ld\n", -(long)((unsigned long) ~0 >> 1));
  printf("signed long max   = %ld\n",  (long)((unsigned long) ~0 >> 1));

  printf("\n");

  printf("Unsigned Types(%s)\n\n", __func__);
  printf("unsigned char max   = %u\n",  (unsigned char)  ~0);
  printf("unsigned short max  = %u\n",  (unsigned short) ~0);
  printf("unsigned int max   = %u\n",   (unsigned int)   ~0);
  printf("unsigned long max   = %lu\n", (unsigned long)  ~0);
}


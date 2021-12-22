#include <stdio.h>

/*
page 36
*/

unsigned long int next = 1;

// 为 custom_rand 设置种子数
void custom_srand(unsigned int seed){
  next = seed;
}

// 返回 0-327667 的伪随机数
int custom_rand(){
  next = next * 1103515245 + 12345;
  return (unsigned int)(next / 65536) % 32768;
}

int main(int argc, char const *argv[])
{
  printf("%d", custom_rand());
  return 0;
}

#include <stdio.h>

int main(int argc, char const *argv[])
{
  int x = 0x66;
  int y = 0x39;

  printf("%x\n", x & y);
  printf("%x\n", x | y);
  printf("%x\n", ~x | ~y);
  printf("%x\n", ~x | !y);


  return 0;
}

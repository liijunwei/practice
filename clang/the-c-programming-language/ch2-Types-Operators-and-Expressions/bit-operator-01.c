#include <stdio.h>

/*
page 39

*/

int main(int argc, char const *argv[]) {

  // 77 -> 0100 1101
  // 45 -> 0010 1101
  printf("按位与\n");
  printf("%d\n", 77 & 45);
  printf("%d\n", 77 & 43);
  printf("%d\n", 77 & 41);
  printf("%d\n", 77 & 39);

  printf("按位或\n");
  printf("%d\n", 77 | 45);
  printf("%d\n", 77 | 43);
  printf("%d\n", 77 | 41);
  printf("%d\n", 77 | 39);

  printf("按位异或\n");
  printf("%d\n", 77 ^ 45);
  printf("%d\n", 77 ^ 43);
  printf("%d\n", 77 ^ 41);
  printf("%d\n", 77 ^ 39);

  printf("左移\n");
  printf("%d\n", 1 << 0);
  printf("%d\n", 1 << 1);
  printf("%d\n", 1 << 2);
  printf("%d\n", 1 << 3);

  printf("右移\n");
  printf("%d\n", 16 >> 0);
  printf("%d\n", 16 >> 1);
  printf("%d\n", 16 >> 2);
  printf("%d\n", 16 >> 3);

  printf("按位求反\n"); // TODO 没明白
  printf("%d\n", ~1);   // 0000 0001
  printf("%d\n", ~7);   // 0000 0111
  printf("%d\n", ~16);  // 0001 0000
  printf("%d\n", ~8);   // 0000 1000

  return 0;
}
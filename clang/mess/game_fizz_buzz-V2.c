#include <stdio.h>

/*
http://wiki.c2.com/?FizzBuzzTest

高手
*/

int main() {
  for (int i = 1; i <= 100; i++) {
    if (i % 3 == 0)
      printf("Fizz");
    if (i % 5 == 0)
      printf("Buzz");
    if (i % 3 != 0 && i % 5 != 0)
      printf("%d", i);

    printf("\n");
  }

  return 0;
}
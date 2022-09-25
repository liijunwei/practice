/*
page 148

Q: 为什么多次调用, 得到的rand值一样?
A: RTFM: "The srand() function sets its argument seed as the seed for a new
sequence of pseudo-random numbers to be returned by rand().  These sequences are
repeatable by calling srand() with the same seed value."
*/

#include <stdio.h>
#include <stdlib.h>

#define frand() ((double)rand() / (RAND_MAX + 1.0))

/* man 3 rand */
int main(int argc, char const *argv[]) {
  srand(1);

  for (int i = 0; i < 5; i++) {
    printf("rand(): %d\n", rand());
  }

  return 0;
}

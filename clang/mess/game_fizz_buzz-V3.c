#include <stdio.h>

/*
http://wiki.c2.com/?FizzBuzzTest

高手
*/

const char *fmt_str[15] = {"FizzBuzz\n",
                           "%d\n",
                           "%d\n",
                           "Fizz\n",
                           "%d\n",
                           "Buzz\n",
                           "Fizz\n"
                           "%d\n",
                           "%d\n",
                           "Fizz\n",
                           "Buzz\n",
                           "%d\n",
                           "Fizz\n",
                           "%d\n",
                           "%d\n"};

// TODO 为什么只能打印到 13???
int main() {
  for (int i = 1; i <= 100; i++) {
    int array_index = i % 15;
    printf(fmt_str[array_index], i);
  }

  return 0;
}
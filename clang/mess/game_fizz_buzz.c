#include <assert.h>
#include <stdio.h>
#include <string.h>

/*
https://en.wikipedia.org/wiki/Fizz_buzz
http://wiki.c2.com/?FizzBuzzTest
https://www.youtube.com/watch?v=CHTep2zQVAc

Players generally sit in a circle. The player designated to go first says the
number "1", and the players then count upwards in turn. However, any number
divisible by three is replaced by the word fizz and any number divisible by five
by the word buzz. Numbers divisible by 15 become fizz buzz. A player who
hesitates or makes a mistake is eliminated from the game. For example, a typical
round of fizz buzz would start as follows: 1, 2, Fizz, 4, Buzz, Fizz, 7, 8,
Fizz, Buzz, 11, Fizz, 13, 14, Fizz Buzz, 16, 17, Fizz, 19, Buzz, Fizz, 22, 23,
Fizz, Buzz, 26, Fizz, 28, 29, Fizz Buzz, 31, 32, Fizz, 34, Buzz, Fizz, ...
*/

char *fizzbuzz(int num);
int is_fizz_buzz(int num);
int is_fizz(int num);
int is_buzz(int num);
char str_buffer[5];

int main() {
  assert(strcmp("1", fizzbuzz(1)) == 0);
  assert(strcmp("2", fizzbuzz(2)) == 0);
  assert(strcmp("fizz", fizzbuzz(3)) == 0);
  assert(strcmp("4", fizzbuzz(4)) == 0);
  assert(strcmp("buzz", fizzbuzz(5)) == 0);
  assert(strcmp("fizz", fizzbuzz(6)) == 0);
  assert(strcmp("fizz buzz", fizzbuzz(15)) == 0);
  printf("ok\n");

  return 0;
}

char *fizzbuzz(int num) {
  if (is_fizz_buzz(num) == 1) {
    return "fizz buzz";
  } else if (is_fizz(num) == 1) {
    return "fizz";
  } else if (is_buzz(num) == 1) {
    return "buzz";
  } else {
    sprintf(str_buffer, "%d", num);
    return str_buffer;
  }
}

int is_fizz_buzz(int num) {
  if (is_fizz(num) && is_buzz(num)) {
    return 1;
  } else {
    return 0;
  }
}

int is_fizz(int num) {
  if (num % 3 == 0) {
    return 1;
  } else {
    return 0;
  }
}

int is_buzz(int num) {
  if (num % 5 == 0) {
    return 1;
  } else {
    return 0;
  }
}

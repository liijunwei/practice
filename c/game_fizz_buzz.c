#include <stdio.h>


/*
https://en.wikipedia.org/wiki/Fizz_buzz
http://wiki.c2.com/?FizzBuzzTest
https://www.youtube.com/watch?v=CHTep2zQVAc

Players generally sit in a circle. The player designated to go first says the number "1", and the players then count upwards in turn. However, any number divisible by three is replaced by the word fizz and any number divisible by five by the word buzz. Numbers divisible by 15 become fizz buzz. A player who hesitates or makes a mistake is eliminated from the game.
For example, a typical round of fizz buzz would start as follows:
1, 2, Fizz, 4, Buzz, Fizz, 7, 8, Fizz, Buzz, 11, Fizz, 13, 14, Fizz Buzz, 16, 17, Fizz, 19, Buzz, Fizz, 22, 23, Fizz, Buzz, 26, Fizz, 28, 29, Fizz Buzz, 31, 32, Fizz, 34, Buzz, Fizz, ...
*/

void fizzbuzz(int num);
int is_fizz_buzz(int num);
int is_fizz(int num);
int is_buzz(int num);

int main(){
  for(int i = 1; i <= 100; i++){
    fizzbuzz(i);
  }

  return 0;
}

void fizzbuzz(int num){
  if(is_fizz_buzz(num) == 1){
    printf("%d fizz buzz\n", num);
  }
  else if(is_fizz(num) == 1){
    printf("%d fizz\n", num);
  }
  else if(is_buzz(num) == 1){
    printf("%d buzz\n", num);
  }
  else{
    printf("%d %d\n", num, num);
  }
}

int is_fizz_buzz(int num){
  if(is_fizz(num) && is_buzz(num)){
    return 1;
  }
  else{
    return 0;
  }
}

int is_fizz(int num){
  if(num % 3 == 0){
    return 1;
  }
  else{
    return 0;
  }
}

int is_buzz(int num){
  if(num % 5 == 0){
    return 1;
  }
  else{
    return 0;
  }
}

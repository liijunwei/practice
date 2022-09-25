#include <stdio.h>

int main() {
  printf("please enter a char: ");
  char c = getchar();
  printf("The input char is: %c\n", c);
  printf("\n");

  putchar(c);
}
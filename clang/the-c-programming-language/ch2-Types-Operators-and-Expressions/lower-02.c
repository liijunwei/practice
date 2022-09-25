#include <ctype.h>
#include <stdio.h>

// page 34

int main(int argc, char const *argv[]) {
  printf("isdight: %d\n", isdigit('a'));
  printf("isdight: %d\n", isdigit(1));
  printf("isdight: %d\n", isdigit(2));

  printf("tolower: %c\n", tolower('A'));
  printf("\n");
  printf("tolower: %c\n", tolower('B'));
  printf("tolower: %c\n", tolower('M'));
  printf("tolower: %c\n", tolower('N'));
  printf("tolower: %c\n", tolower('Y'));
  printf("tolower: %c\n", tolower('Z'));
  printf("tolower: %c\n", tolower('g'));

  return 0;
}

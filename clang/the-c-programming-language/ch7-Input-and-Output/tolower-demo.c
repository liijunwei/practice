/*
page 134

*/

#include <ctype.h>
#include <stdio.h>

// gcc -g ch7-Input-and-Output/tolower-demo.c && echo "THE C PROGRAMING
// LANGUAGE" | ./a.out
int main(int argc, char const *argv[]) {
  int c;

  while ((c = getchar()) != EOF) {
    putchar(tolower(c));
  }

  return 0;
}

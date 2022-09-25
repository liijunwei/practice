#include <stdio.h>

// page 34

int custom_lower(int c) {
  if (c >= 'A' && c <= 'Z') {
    return (c + ('a' - 'A'));
  } else {
    return c;
  }
}

int main(int argc, char const *argv[]) {
  printf("number: %d\n", 'a');
  printf("number: %d\n", 'A');
  printf("number: %d\n", 'a' - 'A');

  printf("lower: %c\n", custom_lower('A'));
  printf("lower: %c\n", custom_lower('B'));
  printf("lower: %c\n", custom_lower('C'));
  printf("lower: %c\n", custom_lower('X'));
  printf("lower: %c\n", custom_lower('Y'));
  printf("lower: %c\n", custom_lower('Z'));

  return 0;
}

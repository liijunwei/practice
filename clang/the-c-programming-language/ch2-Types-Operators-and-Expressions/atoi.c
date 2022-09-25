#include <stdio.h>

// page 33

int custom_atoi(const char s[]) {
  int i;
  int n = 0;

  for (i = 0; s[i] >= '0' && s[i] <= '9'; i++) {
    n = (10 * n) + (s[i] - '0');
  }

  return n;
}

int main(int argc, char const *argv[]) {
  char *num = "9527";
  printf("number: %d\n", custom_atoi(num));

  return 0;
}

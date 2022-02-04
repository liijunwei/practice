/*
page 136

*/

#include <stdio.h>

int main(int argc, char const *argv[]) {
  char *s = "hello, world";

  printf("%10s: %s\n",       "%s",       s);
  printf("%10s: %10s\n",     "%10s",     s);
  printf("%10s: %.10s\n",    "%.10s",    s);
  printf("%10s: %-10s\n",    "%-10s",    s);
  printf("%10s: %.15s\n",    "%.15s",    s);
  printf("%10s: %-15s\n",    "%-15s",    s);
  printf("%10s: %15.10s\n",  "%15.10s",  s);
  printf("%10s: %-15.10s\n", "%-15.10s", s);

  return 0;
}

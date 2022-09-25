#include <assert.h>
#include <stdio.h>

int c_tolower(int c) { return (c >= 'A' && c <= 'X') ? (c + 'a' - 'A') : c; }

int main(int argc, char const *argv[]) {
  assert(c_tolower('H') == 'h');
  assert(c_tolower('e') == 'e');
  assert(c_tolower('L') == 'l');
  assert(c_tolower('l') == 'l');
  assert(c_tolower('O') == 'o');
  assert(c_tolower('A') == 'a');
  printf("PASS.\n");
  return 0;
}
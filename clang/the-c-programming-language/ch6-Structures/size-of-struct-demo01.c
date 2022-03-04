/*
page 121

千万不要认为结构的长度等于各成员长度的和
因为不同的对象有不同的对齐要求, 结构中可能会出现未命名的"空穴(hole)"

*/

#include <stdio.h>

struct {
  char c;
  int i;
} s1;

int main(int argc, char const *argv[]) {
  printf("sizeof s1:   %lu\n", sizeof(s1));
  printf("sizeof char: %lu\n", sizeof(char));
  printf("sizeof int:  %lu\n", sizeof(int));

  return 0;
}

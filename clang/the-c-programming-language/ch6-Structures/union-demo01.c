/*
page 129

TODO
第8章的存储分配程序将说明如何使用联合来强制一个变量在特定类型的存储边界上对齐

*/

#include <stdio.h>

#define NSYM 10

union u_tag {
  int ival;
  float fval;
  char *svar;
} u;

struct {
  char *name;
  int flags;
  int utype;
  union {
    int ival;
    float fval;
    char *svar;
  } u;
} systab[NSYM];

int main(int argc, char const *argv[]) { return 0; }

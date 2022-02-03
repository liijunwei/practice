/*
page 129

TODO

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

int main(int argc, char const *argv[]) {


  return 0;
}

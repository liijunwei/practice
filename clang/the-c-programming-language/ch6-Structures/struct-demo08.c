/*
page 122

自引用结构的变体
*/

#include <stdio.h>

struct t {
  struct s *p;
};

struct s {
  struct t *q;
};

int main(int argc, char const *argv[]) { return 0; }

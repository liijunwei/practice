#include <stdio.h>
// https://www.runoob.com/w3cnote/c-bool-true-false.html
#include <stdbool.h>
#include <string.h>

// gcc foobar/stdbool-demo.c && ./a.out demo
// gcc foobar/stdbool-demo.c && ./a.out demo1
// run foobar/stdbool-demo.c demo
// run foobar/stdbool-demo.c demo1

int main(int argc, char const *argv[]) {
  bool flag = strcmp(argv[1], "demo") == 0 ? true : false;

  if (flag) {
    printf("flag is true");
  } else {
    printf("flag is false");
  }

  return 0;
}

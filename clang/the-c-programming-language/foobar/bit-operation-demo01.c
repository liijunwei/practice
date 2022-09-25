/*
page 105

*/

#include <ctype.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#include "../common-utils/print-error.c"

#define NUMERIC 1 /* -n numeric sort                     */
#define DECR 2    /* -r sort in decreasing order         */
#define FOLD 4    /* -f fold upper and lower cases       */
#define DIR 8     /* -d dir order                        */

#define LINES 100 /* -n max number of lines to be sorted */

void readargs(int argc, char const *argv[]);

static char option = 0;

// gcc foobar/bit-operation-demo01.c && ./a.out -n # 1
// gcc foobar/bit-operation-demo01.c && ./a.out -r # 2
// gcc foobar/bit-operation-demo01.c && ./a.out -f # 4
// gcc foobar/bit-operation-demo01.c && ./a.out -d # 8
// gcc foobar/bit-operation-demo01.c && ./a.out -nr # 3
// gcc foobar/bit-operation-demo01.c && ./a.out -nf # 5
// gcc foobar/bit-operation-demo01.c && ./a.out -nd # 9
// gcc foobar/bit-operation-demo01.c && ./a.out -fd # 12
// gcc foobar/bit-operation-demo01.c && ./a.out -fr # 6
// gcc foobar/bit-operation-demo01.c && ./a.out -nfd # 13
// gcc foobar/bit-operation-demo01.c && ./a.out -nrfd # 15
int main(int argc, char const *argv[]) {
  readargs(argc, argv);
  printf("option is: %d\n", option);

  return 0;
}

void readargs(int argc, char const *argv[]) {
  int c;

  while (--argc > 0 && (c = (*++argv)[0]) == '-' || c == '+') {
    if (c == '-' && !isdigit(*(argv[0] + 1))) {
      while ((c = *++argv[0]) > 0) {
        switch (c) {
        case 'n':
          option |= NUMERIC;
          break;
        case 'r':
          option |= DECR;
          break;
        case 'f':
          option |= FOLD;
          break;
        case 'd':
          option |= DIR;
          break;
        default:
          error("Error: unsopported option");
          break;
        }
      }
    }
  }
}

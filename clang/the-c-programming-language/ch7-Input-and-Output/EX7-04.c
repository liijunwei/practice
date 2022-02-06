/*
page 140

类似于 EX7-03的minscanf, 编写一个scanf函数的简化版本

*/

#include <stdio.h>
#include <stdarg.h>
#include <ctype.h>

#define LOCALFMT 100

void minscanf(char *fmt, ...);

// gcc -g ch7-Input-and-Output/EX7-04.c && echo 20 "Monkey-D-Luffy" | ./a.out
int main(int argc, char const *argv[]) {
  int number;
  printf("Please enter a number: \n");
  minscanf("%d", &number);
  printf("You entered: %d\n", number);

  char name[100];
  printf("Please enter a name: \n");
  minscanf("%s", name);
  printf("You entered: %s\n", name);

  return 0;
}

/* minimal scanf with variable argument list */
void minscanf(char *fmt, ...) {
  va_list ap;
  char *p;
  char *sval;
  char localfmt[LOCALFMT];
  int c;
  int i = 0; /* index for localfmt array */
  int *ival;
  unsigned int *uval;
  double *dval;

  va_start(ap, fmt);

  for (p = fmt; *p; p++) {
    if (*p != '%') {
      localfmt[i++] = *p;
      continue;
    }

    localfmt[i++] = '%';

    while (*(p+1) && !isalpha(*(p+1))) {
      localfmt[i++] = *++p; /* collect chars */
    }

    localfmt[i++] = *(p+1);
    localfmt[i] = '\0';

    switch (*++p) {
      case 'd':
      case 'i':
        ival = va_arg(ap, int *);
        scanf(localfmt, ival);
        break;
      case 'x':
      case 'X':
      case 'u':
      case 'o':
        uval = va_arg(ap, unsigned int *);
        scanf(localfmt, uval);
        break;
      case 'f':
        dval = va_arg(ap, double *);
        scanf(localfmt, dval);
        break;
      case 's':
        sval = va_arg(ap, char *);
        scanf(localfmt, sval);
        break;
      default:
        putchar(*p);
        break;
    }

    i = 0; /* reset index */
  }

  va_end(ap);
}

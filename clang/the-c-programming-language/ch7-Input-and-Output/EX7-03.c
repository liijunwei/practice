/*
page 137

改写minprintf函数, 使他能完成printf的更多功能

*/

#include <stdio.h>
#include <stdarg.h>
#include <ctype.h>

#define LOCALFMT 100

void minprintf(char *fmt, ...);

int main(int argc, char const *argv[]) {
  minprintf("two person laughing(d) | %d\n", 466666);
  minprintf("two person laughing(i) | %i\n", 466666);
  minprintf("unsigned int(x)        | %x\n", 998);
  minprintf("unsigned int(X)        | %X\n", 776);
  minprintf("unsigned int(u)        | %u\n", 666);
  minprintf("unsigned int(o)        | %o\n", 111);
  minprintf("constant               | %f\n", 3.14159265354);
  minprintf("say hello to           | %s\n", "Jef");
  minprintf("random string          | %abc\n", "laksdjflkj");

  return 0;
}

void minprintf(char *fmt, ...) {
  va_list ap;
  char *p;
  char *sval;
  char localfmt[LOCALFMT];
  int i;
  int ival;
  unsigned int uval;
  double dval;

  va_start(ap, fmt);

  for (p = fmt; *p; p++) {
    if (*p != '%') {
      putchar(*p);
      continue;
    }

    i = 0;
    localfmt[i++] = '%';

    while (*(p+1) && !isalpha(*(p+1))) {
      localfmt[i++] = *++p; /* collect chars */
    }

    localfmt[i++] = *(p+1);
    localfmt[i] = '\0';

    switch (*++p) {
      case 'd':
      case 'i':
        ival = va_arg(ap, int);
        printf(localfmt, ival);
        break;
      case 'x':
      case 'X':
      case 'u':
      case 'o':
        uval = va_arg(ap, unsigned int);
        printf(localfmt, uval);
        break;
      case 'f':
        dval = va_arg(ap, double);
        printf(localfmt, dval);
        break;
      case 's':
        sval = va_arg(ap, char *);
        printf(localfmt, sval);
        break;
      default:
        putchar(*p);
        break;
    }
  }

  va_end(ap);
}

#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <ctype.h>

#include "../common-utils/getline.c"
#include "../common-utils/readlines.c"
#include "../common-utils/numcmp.c"

/*
page 105

增加-f选项, 使得排序过程不考虑字母大小写之间的区别
例如: 比较a和A时认为他们相等

*/

#define NUMERIC 1
#define DECR    2   /* sort in decreasing order */
#define FOLD    4   /* fold upper and lower cases */
#define LINES   100 /* max number of lines to be sorted */

int charcmp(char *, char *);
void custom_qsort(char *v[], int left, int right, int (*comp)(void *, void *));
void printlines(char *lineptr[], int nlines, int order);

static char option = 0;

/* sort input lines */
int main(int argc, char const *argv[])
{

  return 0;
}

void printlines(char *lineptr[], int nlines, int decr) {
  int i;

  if(decr) {
    for(i = nlines - 1; i >= 0; i--) {
      printf("%s\n", lineptr[i]);
    }
  } else {
    for(i = 0; i < nlines; i++) {
      printf("%s\n", lineptr[i]);
    }
  }
}

/* return < 0 if s < t  */
/* return   0 if s == t */
/* return > 0 if s > t  */
int charcmp(char *s, char *t) {
  for(; tolower(*s) == tolower(*t); s++, t++) {
    if(*s == '\0') {
      return 0;
    }
  }

  return tolower(*s) - tolower(*t);
}




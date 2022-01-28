/*
page 105

增加选项-d(代表目录顺序)
该选项表明 只对字母/数字/和空格进行比较
要保证该选项可以和-f组合在一起使用

TODO 没懂...
*/

#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <ctype.h>

#include "../common-utils/getline.c"
#include "../common-utils/readlines.c"
#include "../common-utils/numcmp.c"

#define NUMERIC 1
#define DECR    2   /* sort in decreasing order         */
#define FOLD    4   /* fold upper and lower cases       */
#define DIR     8   /* dir order                        */
#define LINES   100 /* max number of lines to be sorted */

int charcmp(char *, char *);
void custom_qsort(void *v[], int left, int right, int (*comp)(void *, void *));
void printlines(char *lineptr[], int nlines, int order);

static char option = 0;

// cat ch5-Pointers-and-Arrays/sample-input.md | sort
// gcc -g ch5-Pointers-and-Arrays/EX5-16.c && cat ch5-Pointers-and-Arrays/sample-input.md | ./a.out
// gcc -g ch5-Pointers-and-Arrays/EX5-16.c && cat ch5-Pointers-and-Arrays/sample-input.md | ./a.out -fnr

/* sort input lines */
int main(int argc, char const *argv[])
{
  char *lineptr[LINES];
  int nlines;
  int c;
  int rc = 0;

  while(--argc > 0 && (*++argv)[0] == '-') {
    while((c = *++argv[0]) > 0) {
      switch(c) {
        case 'd':
          option |= DIR;
          break;
        case 'f':
          option |= FOLD;
          break;
        case 'n':
          option |= NUMERIC;
          break;
        case 'r':
          option |= DECR;
          break;
        default:
          printf("Sort: illegal option %c\n", c);
          argc = 1;
          rc = -1;
          break;
      }
    }
  }

  if(argc) {
    printf("Usage: sort -dfnr\n");
  } else {
    if((nlines = readlines(lineptr, LINES)) > 0) {
      if(option & NUMERIC) {
        custom_qsort((void **)lineptr, 0, nlines - 1, (int(*)(void *, void *)) numcmp);
      } else {
        custom_qsort((void **)lineptr, 0, nlines - 1, (int(*)(void *, void *)) strcmp);
      }

      printlines(lineptr, nlines, option & DECR);
    } else {
      printf("input too big to sort\n");
      rc = -1;
    }
  }

  return rc;
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
  char a;
  char b;

  int fold = (option & FOLD) ? 1 : 0;
  int dir  = (option & DIR)  ? 1 : 0;

  do {
    if(dir) {
      while(!isalnum(*s) && *s != ' ' && *s != '\0') {
        s++;
      }

      while(!isalnum(*t) && *t != ' ' && *t != '\0') {
        s++;
      }
    }

    a = fold ? tolower(*s) : *s;
    s++;
    b = fold ? tolower(*t) : *t;
    t++;

    if(a == b && a == '\0') {
      return 0;
    }
  } while(a == b);

  return a - b;
}

void swap(void *v[], int i, int j){
  void *temp = v[i];
  v[i] = v[j];
  v[j] = temp;
}

void custom_qsort(void *v[], int left, int right, int (*comp)(void *, void *)) {
  int i;
  int last;

  if(left >= right){ // 如果数组元素的个数小于2, 则返回
    return;
  }

  swap(v, left, (left + right) / 2);
  last = left;

  for(i = left + 1; i <= right; i++){
    if((*comp)(v[i], v[left]) < 0){ // asc
      swap(v, ++last, i);
    }
  }

  swap(v, left, last);
  custom_qsort(v, left, last - 1, comp);
  custom_qsort(v, last + 1, right, comp);
}


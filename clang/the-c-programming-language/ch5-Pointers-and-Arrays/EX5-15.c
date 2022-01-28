#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <ctype.h>

#include "../common-utils/getline.c"

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
int numcmp(char *, char *);
int readlines(char *lineptr[], int maxlines);
void custom_qsort(char *v[], int left, int right, int (*comp)(void *, void *));
void printlines(char *lineptr[], int nlines, int order);

static char option = 0;

/* sort input lines */
int main(int argc, char const *argv[])
{

  return 0;
}


#define MAXLEN 1000
#define ALLOCSIZE 10000
static char allocbuf[ALLOCSIZE];
static char *allocp = allocbuf;

// 返回指向n个字符的指针(ch5-Pointers-and-Arrays/allocbuf-afree-demo.c)
char *alloc(int n) {
  if(allocbuf + ALLOCSIZE - allocp >= n) {
    allocp += n;
    return allocp - n;
  } else {
    printf("memory not enough(%d > %d)\n", n, ALLOCSIZE);
    return NULL;
  }
}

int readlines(char *lineptr[], int maxlines) {
  int len;
  int nlines;
  char *p;
  char line[MAXLEN];

  nlines = 0;

  while((len = custom_getline(line, MAXLEN)) > 0) {
    if(nlines >= maxlines || (p = alloc(len)) == NULL) {
      return -1;
    } else {
      line[len - 1] = '\0'; /* 把 '\n' 替换为 '\0' */
      strcpy(p, line); /* char* strcpy( char* dest, const char* src ); */
      lineptr[nlines++] = p;
    }
  }

  return nlines;
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

int numcmp(char *s1, char *s2) {
  double v1;
  double v2;

  v1 = atof(s1);
  v2 = atof(s2);

  if(v1 < v2) {
    return -1;
  } else if (v1 > v2) {
    return 1;
  } else {
    return 0;
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




#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#include "../common-utils/getline.c"

/*
page 105

修改排序程序, 是他能处理-r参数
该参数表明 以递减方式排序, 要保证-n和-r能够组合在一起使用

*/

#define NUMERIC 1   /* numeric sort */
#define DECR    2   /* sort in decreasing order */
#define LINES   100 /* max number of lines to be sorted */

int numcmp(char *, char *);
int readlines(char *lineptr[], int nlines);
void custom_qsort(void *lineptr[], int left, int right, int (*comp)(void *, void *));
void printlines(char *lineptr[], int nlines, int decr);

static char option = 0;

// gcc -g ch5-Pointers-and-Arrays/EX5-14.c
//
// cat ch5-Pointers-and-Arrays/tmp.md | ./a.out
// cat ch5-Pointers-and-Arrays/tmp.md | ./a.out -n
// cat ch5-Pointers-and-Arrays/tmp.md | ./a.out -r
// cat ch5-Pointers-and-Arrays/tmp.md | ./a.out -nr
int main(int argc, char const *argv[])
{
  char *lineptr[LINES];
  int nlines;
  int c;
  int rc = 0;

  while(--argc > 0 && (*++argv)[0] == '-') {
    while((c = *++argv[0]) > 0) {
      switch(c) {
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
    printf("Usage: sort -nr\n");
  } else {
    if((nlines = readlines(lineptr, LINES)) > 0) {
      if(option & NUMERIC) {
        qsort((void **)lineptr, 0, nlines - 1, (int(*)(void *, void *)) numcmp);
      } else {
        qsort((void **)lineptr, 0, nlines - 1, (int(*)(void *, void *)) strcmp);
      }

      printlines(lineptr, nlines, option & DECR);
    } else {
      printf("input too big to sort\n");
      rc = -1;
    }
  }

  return rc;
}

#define MAXLEN 1000 // 每个输入行的最大长度
#define ALLOCSIZE 10000 // 可用空间大小
static char allocbuf[ALLOCSIZE]; // alloc使用的存储区
static char *allocp = allocbuf;  // 下一空闲位置

// 返回指向n个字符的指针(ch5-Pointers-and-Arrays/allocbuf-afree-demo.c)
char *alloc(int n){
  if(allocbuf + ALLOCSIZE - allocp >= n){
    allocp += n;
    return allocp - n; // 明白了: allocp 这个指针在向右移动了n个单位后, 到了新的位置, 返回一个新的位置(指针), 表示分配好的这n个单位内存空间
  } else {
    printf("memory not enough(%d > %d)\n", n, ALLOCSIZE);
    return NULL; // c语言保证, 0永远不是有效的数据地址, 因此可以用做没有足够空间的标识
  }
}

int readlines(char *lineptr[], int maxlines){
  int len;
  int nlines;
  char *p;
  char line[MAXLEN];

  nlines = 0;

  while((len = custom_getline(line, MAXLEN)) > 0){
    if(nlines >= maxlines || (p = alloc(len)) == NULL){
      return -1;
    } else {
      line[len - 1] = '\0'; // 把 '\n' 替换为 '\0'
      strcpy(p, line); // char* strcpy( char* dest, const char* src );
      lineptr[nlines++] = p;
    }
  }

  return nlines;
}


void printlines(char *lineptr[], int nlines, int decr){
  int i;

  if(decr) {
    for(i = nlines - 1; i >= 0; i--){
      printf("%s\n", lineptr[i]);
    }
  } else {
    for(i = 0; i < nlines; i++){
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

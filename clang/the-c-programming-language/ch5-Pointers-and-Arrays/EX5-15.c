#include <stdio.h>
#include <string.h>
#include <ctype.h>

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
int readlines(char *, char *);
void custom_qsort(char *v[], int left, int right, int (*comp)(void *, void *));
void printlines(char *lineptr[], int nlines, int order);

static char option = 0;

/* sort input lines */
int main(int argc, char const *argv[])
{

  return 0;
}

#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#include "../common-utils/getline.c"

/*
page 102

编写程序tail, 将其输入中的最后n行打印出来
默认情况下, n的值为10, 但可通过一个可选参数改变n的值, 因此, 命令 `tail -n` 将打印其输入的最后n行
无论输入或n的值是否合理, 该程序都应该能正常运行

编写的程序要重复利用存储空间, 输入行的存储方式应该通5.6节中排序程序的存储方式一样, 而不采用固定长度的二维数组

*/

#define DEFAULT_LINES 10  // default number of lines to print
#define LINES         100 // max number of lines to print
#define MAXLEN        10  // max length of an input line

void error(char *s);

// print last n lines of the input
int main(int argc, char const *argv[])
{
  char *p;
  char *buf;
  char *bufend;
  char line[MAXLEN];
  char *lineptr[LINES];
  int first;
  int i;
  int last;
  int len;
  int n;
  int nlines;

  if(argc == 1) {
    n = DEFAULT_LINES;
  } else if(argc == 2 && (*++argv)[0] == '-') {
    n = atoi(argv[1] + 1);
  } else {
    error("usage: tail [-n]");
  }

  if(n < 1 || n > LINES) {
    n = LINES;
  }

  for(i = 0; i < LINES; i++) {
    lineptr[i] = NULL;
  }

  if((p = buf = malloc(LINES * MAXLEN)) == NULL) {
    error("tail: cannot allocate buf");
  }

  bufend = buf + LINES * MAXLEN;
  last = 0;
  nlines = 0;

  while((len = custom_getline(line, MAXLEN)) > 0) {
    if(p + len + 1 >= bufend) {
      p = buf;
    }

    lineptr[last] = p;
    strcpy(lineptr[last], line);

    if(++last >= LINES) {
      last = 0;
    }

    p += len + 1;
    nlines++;
  }

  if(n > nlines) {
    n = nlines;
  }

  first = last - n;

  if(first < 0) {
    first += LINES;
  }

  for(i = first; n-- > 0; i = (i + 1) % LINES) {
    printf("%s", lineptr[i]);
  }

  return 0;
}


// print error message and quit
void error(char *s){
  printf("%s\n", s);
  exit(1);
}


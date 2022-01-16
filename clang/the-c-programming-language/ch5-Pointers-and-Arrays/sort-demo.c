#include <stdio.h>
#include <string.h>

/*
page 93

*/

#define MAXLINES 5000 // 进行排序的最大文本行数

char *lineptr[MAXLINES];
int readlines(char *lineptr[], int nlines);
void writelines(char *lineptr[], int nlines);
void qsort(char *lineptr[], int left, int right);

// 对输入的文本行进行排序
int main(int argc, char const *argv[])
{
  int nlines;

  if((nlines = readlines(lineptr, MAXLINES)) >= 0){
    qsort(lineptr, 0, nlines - 1);
    writelines(lineptr, nlines);
    return 0;
  } else {
    printf("error: input too big to sort\n");
    return 1;
  }
}

#define MAXLEN 1000 // 每个输入行的最大长度
char *alloc(int);

int custom_getline(char s[], int max){
  int c;
  int i = 0;

  while(--max > 0 && (c = getchar()) != EOF && c != '\n'){
    s[i] = c;
    i++;
  }

  if(c == '\n'){
    s[i] = c;
    ++i;
  }

  s[i] = '\0';

  return i;
}

// 读取该输入行
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
      line[len - 1] = '\0';
      strcpy(p, line);
      lineptr[nlines++] = p;
    }
  }

  return nlines;
}

void writelines(char *lineptr[], int nlines){
  int i;

  for(i = 0; i < nlines; i++){
    printf("%s\n", lineptr[i]);
  }
}


#include <stdio.h>
#include <string.h>

/*
page 95

ch5-Pointers-and-Arrays/sort-demo.c

重写函数readlines, 将输入的文本行存储到由main函数提供的一个数组中, 而不是存储到调用alloc分配的存储空间中;
改函数的运行速度比改写之前快多少?

*/

#define MAXLINES 5000 // 进行排序的最大文本行数

char *lineptr[MAXLINES];
int readlines(char *lineptr[], char *linestor, int maxlines);
void printlines(char *lineptr[], int nlines);
void qsort(char *lineptr[], int left, int right);

// gcc -g ch5-Pointers-and-Arrays/EX5-07.c && ./a.out < ch5-Pointers-and-Arrays/tmp.md
// 对输入的文本行进行排序
int main(int argc, char const *argv[])
{
  int nlines;
  char linestor[100];

  if((nlines = readlines(lineptr, linestor, MAXLINES)) >= 0){
    qsort(lineptr, 0, nlines - 1);
    printlines(lineptr, nlines);
    return 0;
  } else {
    printf("error: input too big to sort\n");
    return 1;
  }
}

#define MAXLEN 1000 // 每个输入行的最大长度
#define MAXSTOR 5000 // size of available storage space

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

int readlines(char *lineptr[], char *linestor, int maxlines){
  int len;
  int nlines;
  char *p = linestor;
  char line[MAXLEN];
  char *linestop = linestor + MAXSTOR;

  nlines = 0;

  while((len = custom_getline(line, MAXLEN)) > 0){
    if(nlines >= maxlines || (p + len > linestop)){
      return -1;
    } else {
      line[len - 1] = '\0'; // 把 '\n' 替换为 '\0'
      strcpy(p, line); // char* strcpy( char* dest, const char* src );
      lineptr[nlines++] = p;
      p += len;
    }
  }

  return nlines;
}

void printlines(char *lineptr[], int nlines){
  int i;

  for(i = 0; i < nlines; i++){
    printf("%s\n", lineptr[i]);
  }
}

void swap(char *v[], int i, int j){
  char *temp = v[i];
  v[i] = v[j];
  v[j] = temp;
}

// TODO not clear
void qsort(char *v[], int left, int right){
  int i;
  int last;

  if(left >= right){ // 如果数组元素的个数小于2, 则返回
    return;
  }

  swap(v, left, (left + right) / 2);
  last = left;

  for(i = left + 1; i <= right; i++){
    if(strcmp(v[i], v[left]) < 0){ // asc
      swap(v, ++last, i);
    }
  }

  swap(v, left, last);
  qsort(v, left, last - 1);
  qsort(v, last + 1, right);
}



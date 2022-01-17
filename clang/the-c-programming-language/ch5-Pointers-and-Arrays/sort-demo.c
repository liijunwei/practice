#include <stdio.h>
#include <string.h>

/*
page 93

*/

#define MAXLINES 5000 // 进行排序的最大文本行数

char *lineptr[MAXLINES];
int readlines(char *lineptr[], int nlines);
void printlines(char *lineptr[], int nlines);
void qsort(char *lineptr[], int left, int right);

// TODO 没全看懂...
// 对输入的文本行进行排序
int main(int argc, char const *argv[])
{
  int nlines;

  if((nlines = readlines(lineptr, MAXLINES)) >= 0){
    qsort(lineptr, 0, nlines - 1);
    printlines(lineptr, nlines);
    return 0;
  } else {
    printf("error: input too big to sort\n");
    return 1;
  }
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

void qsort(char *v[], int left, int right){
  int i;
  int last;

  if(left >= right){ // 如果数组元素的个数小于2, 则返回
    return;
  }

  swap(v, left, (left + right) / 2);
  last = left;

  for(i = left + 1; i <= right; i++){
    if(strcmp(v[i], v[left]) < 0){
      swap(v, ++last, i);
    }
  }

  swap(v, left, last);
  qsort(v, left, last - 1);
  qsort(v, last + 1, right);
}



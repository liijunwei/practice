#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#include "../common-utils/getline.c"

/*
page 102

修改前面的排序函数, 在给定可选参数-n的情况下, 该函数将按数值大小而非字典顺序对输入进行排序

排序程序通常包含3部分:
  判断任何两个对象之间次序的比较操作
  颠倒对象次序的交换操作
  用于比较和交换对象直到所有对象都按正确次序排列的排序算法

由于排序算法和比较/交换操作无关, 因此, 通过在排序算法中调用不同的比较和交换函数, 便可以实现按照不同的标准排序

TODO ch5-Pointers-and-Arrays/sort-demo2.c:42:83: warning: pointer type mismatch ('int (*)(char *, char *)' and 'int (*)(const char *, const char *)') [-Wpointer-type-mismatch]

*/

#define MAXLINES 5000 // 进行排序的最大文本行数

char *lineptr[MAXLINES];
int readlines(char *lineptr[], int nlines);
void printlines(char *lineptr[], int nlines);
void custom_qsort(void *lineptr[], int left, int right, int (*comp)(void *, void *));
int numcmp(char *s1, char *s2);

// gcc -g ch5-Pointers-and-Arrays/sort-demo2.c && cat ch5-Pointers-and-Arrays/sample-input.md | ./a.out
// gcc -g ch5-Pointers-and-Arrays/sort-demo2.c && cat ch5-Pointers-and-Arrays/sample-input.md | ./a.out -n
// 对输入的文本行进行排序
int main(int argc, char const *argv[])
{
  int nlines;
  int numeric = 0; /* 若进行数值排序, 则numeric的值为1 */

  if(argc > 1 && strcmp(argv[1], "-n") == 0) {
    numeric = 1;
  }

  if((nlines = readlines(lineptr, MAXLINES)) >= 0){
    custom_qsort((void **)lineptr, 0, nlines - 1, (int (*)(void*, void*))(numeric ? numcmp : strcmp));
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

void swap(void *v[], int i, int j){
  void *temp = v[i];
  v[i] = v[j];
  v[j] = temp;
}

// TODO not clear
// 以递增顺序对v[left]...v[right]进行排序
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


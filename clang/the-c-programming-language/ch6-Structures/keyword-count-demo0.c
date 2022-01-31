/*
page 116

统计输入中各个C语言关键字出现的次数
*/

#include <stdio.h>
#include <ctype.h>
#include <string.h>

struct key {
  char *word;
  int count;
} keytab[] = {
  {"auto",     0},
  {"break",    0},
  {"case",     0},
  {"char",     0},
  {"const",    0},
  {"continue", 0},
  {"default",  0},
  {"unsigned", 0},
  {"void",     0},
  {"lolatile", 0},
  {"while",    0},
};

#define NKEYS   11 /* keytab中关键字的个数 */
#define MAXWORD 100

int getword(char *, int);
int binsearch(char *word, struct key tab[], int n);

int main(int argc, char const *argv[])
{
  int n;
  char word[MAXWORD];

  while(getword(word, MAXWORD) != EOF) {
    if(isalpha(word[0])) {
      if((n = binsearch(word, keytab, NKEYS)) >= 0) {
        keytab[n].count++;
      }
    }
  }

  for(n = 0; n < NKEYS; n++) {
    if(keytab[n].count > 0) {
      printf("%4d %s\n", keytab[n].count, keytab[n].word);
    }
  }

  return 0;
}

/* 每调用一次该函数, 将读入一个单词, 并将其复制到名字为该函数的第一个参数的数组中 */
int getword(char *, int) {

}

int binsearch(char *word, struct key tab[], int n) {
  int low = 0;
  int mid;
  int high = n -1;

  int cond;

  while(low <= high){
    mid = (low + high) >> 1;

    if((cond = strcmp(word, tab[mid].word)) < 0) {
      high = mid - 1;
    } else if(cond > 0) {
      low = mid + 1;
    } else {
      return mid;
    }
  }

  return -1;
}
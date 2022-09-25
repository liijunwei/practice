/*
page 116

统计输入中各个C语言关键字出现的次数

isalpha 是否为字母
isalnum 是否为字母或数字
sizeof is an operator
*/

#include <ctype.h>
#include <stdio.h>
#include <string.h>

#include "../common-utils/getch-ungetch.c"

struct key {
  char *word;
  int count;
} keytab[] = {
    {"auto", 0},  {"break", 0},    {"case", 0},    {"char", 0},
    {"const", 0}, {"continue", 0}, {"default", 0}, {"unsigned", 0},
    {"void", 0},  {"lolatile", 0}, {"while", 0},
};

#define NKEYS (sizeof(keytab) / sizeof(struct key)) /* keytab中关键字的个数 */
#define MAXWORD 100

int getword(char *word, int limit);
int binsearch(char *word, struct key tab[], int n);

// bash ch6-Structures/keyword-count-test.sh
int main(int argc, char const *argv[]) {
  int n;
  char word[MAXWORD];

  while (getword(word, MAXWORD) != EOF) {
    if (isalpha(word[0])) {
      if ((n = binsearch(word, keytab, NKEYS)) >= 0) {
        keytab[n].count++;
      }
    }
  }

  for (n = 0; n < NKEYS; n++) {
    if (keytab[n].count > 0) {
      printf("%4d %s\n", keytab[n].count, keytab[n].word);
    }
  }

  return 0;
}

/* TODO 没看懂... */
/* 每调用一次该函数, 将读入一个单词, 并将其复制到 word 中 */
int getword(char *word, int limit) {
  int c;
  char *w = word;

  while (isspace(c = getch()) && c != '\n') {
    ;
  }

  if (c != EOF) {
    *w++ = c; /* * 和 & 的优先级高于算数运算符, 因此这里等价于 (*w) = c; w++ */
  }

  if (!isalpha(c)) {
    *w = '\0';
    return c;
  }

  for (; --limit > 0; w++) {
    if (!isalnum(*w = getch())) {
      ungetch(*w);
      break;
    }
  }

  *w = '\0';
  return word[0];
}

/* 注意: 关键字列表必须按升序存储在keytab里 */
int binsearch(char *word, struct key tab[], int n) {
  int low = 0;
  int mid;
  int high = n - 1;

  int cond;

  while (low <= high) {
    mid = (low + high) >> 1;

    if ((cond = strcmp(word, tab[mid].word)) < 0) {
      high = mid - 1;
    } else if (cond > 0) {
      low = mid + 1;
    } else {
      return mid;
    }
  }

  return -1;
}

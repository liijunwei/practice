/*
page 119

为了进一步说明指结构的指针和结构数组, 重新编写关键字统计程序, 这次采用指针, 不用数组下标

*/

#include <stdio.h>
#include <ctype.h>
#include <string.h>

#include "../common-utils/getch-ungetch.c"

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

#define NKEYS   (sizeof(keytab) / sizeof(struct key)) /* keytab中关键字的个数 */
#define MAXWORD 100

int getword(char *word, int limit);
struct key *binsearch(char *word, struct key *tab, int n);

// bash ch6-Structures/keyword-count-test.sh
int main(int argc, char const *argv[])
{
  char word[MAXWORD];
  struct key *p;

  while(getword(word, MAXWORD) != EOF) {
    if(isalpha(word[0])) {
      if((p = binsearch(word, keytab, NKEYS)) != NULL) {
        p->count++;
      }
    }
  }

  for(p = keytab; p < keytab + NKEYS; p++) {
    if(p->count > 0) {
      printf("%4d %s\n", p->count, p->word);
    }
  }

  return 0;
}

/* 每调用一次该函数, 将读入一个单词, 并将其复制到名字为该函数的第一个参数的数组中 */
int getword(char *word, int limit) {
  int c;
  char *w = word;

  while(isspace(c = getch()) && c != '\n') {
    ;
  }

  if(c != EOF) {
    *w++ = c;
  }

  if(!isalpha(c)) {
    *w = '\0';
    return c;
  }

  for( ; --limit > 0; w++) {
    if(!isalnum(*w = getch())) {
      ungetch(*w);
      break;
    }
  }

  *w = '\0';
  return word[0];
}

struct key *binsearch(char *word, struct key *tab, int n) {
  struct key *low  = &tab[0];
  struct key *high = &tab[n];
  struct key *mid;

  int cond;

  while(low < high) {
    mid = low + (high - low)/2;

    if((cond = strcmp(word, mid->word)) < 0) {
      high = mid;
    } else if(cond > 0) {
      low = mid + 1;
    } else {
      return mid;
    }
  }

  return NULL;
}

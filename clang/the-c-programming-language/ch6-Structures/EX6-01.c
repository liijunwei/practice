/*
page 119

ch6-Structures/keyword-count-demo0.c 里的getword函数不能正确处理
下划线/字符串常量/注释/预处理器控制指令 请编写一个更完善的getword函数

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

// gcc ch6-Structures/EX6-01.c && cat ch6-Structures/keyword-count-sample.md |
// ./a.out
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

int comment() {
  int c;
  while ((c = getch()) != EOF) {
    if (c == '*') {
      if ((c = getch()) == '/') {
        break;
      } else {
        ungetch(c);
      }
    }
  }

  return c;
}

/* TODO 没看懂... */
/* get next word or character from input */
/* 每调用一次该函数, 将读入一个单词,
 * 并将其复制到名字为该函数的第一个参数的数组中 */
int getword(char *word, int limit) {
  int c;
  int d;
  char *w = word;

  while (isspace(c = getch()) && c != '\n') {
    ;
  }

  if (c != EOF) {
    *w++ = c;
  }

  if (isalpha(c) || c == '_' || c == '#') {
    for (; --limit > 0; w++) {
      if (!isalnum(*w = getch()) && *w != '_') {
        ungetch(*w);
        break;
      }
    }
  } else if (c == '\'' || c == '"') {
    for (; --limit > 0; w++) {
      if ((*w = getch()) == '\\') {
        *++w = getch();
      } else if (*w == c) {
        w++;
        break;
      } else if (*w == EOF) {
        break;
      }
    }
  } else if (c == '/') {
    if ((d = getch()) == '#') {
      c = comment();
    } else {
      ungetch(d);
    }
  }

  *w = '\0';
  return c;
}

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

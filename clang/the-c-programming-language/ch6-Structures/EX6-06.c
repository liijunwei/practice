/*
page 127

以本节(6.6 表查找)介绍的函数为基础,
编写一个适合C语言程序使用的#define处理器的简单版本(即无参数的情况)
你会发现getch和ungetch函数非常有用

TODO 没看懂, 怎么测试呢

*/

#include <ctype.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#include "../common-utils/getch-ungetch.c"

struct nlist {
  struct nlist *next; /* pointer to next node       */
  char *name;         /* defined name to substitute */
  char *defn;         /* real text                 */
};

#define HASHSIZE 101
#define MAXWORD 100

unsigned int hash(char *s);
static struct nlist *hashtable[HASHSIZE];
struct nlist *install(char *name, char *defn);
struct nlist *lookup(char *s);

void undef(char *s);
void getdef();
void error(int, int, char *);
int getword(char *, int);
void skipblanks();
void printtable();
void printinfo(int linenum, char *s) { printf("line(%d) -> %s\n", linenum, s); }

/* simple version of #define processor */
int main(int argc, char const *argv[]) {
  // TODO how to test ?

  char word[MAXWORD];
  struct nlist *p;

  install("name", "lijunwei");

  while (getword(word, MAXWORD) != EOF) {
    printinfo(__LINE__, word);
    if (strcmp(word, "#") == 0) { /* begining of directive */
      getdef();
    } else if (!isalpha(word[0])) { /* cannot be defined */
      printf("%s", word);
    } else if ((p = lookup(word)) == NULL) { /* not defined */
      printf("%s", word);
    } else { /* push definition */
      ungets(p->defn);
    }
  }

  printtable();
  return 0;
}

/* get next word or character from input */
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
  }

  *w = '\0';
  return c;
}

/* print error message and skip the rest of the line */
void error(int c, int linenum, char *s) {
  printf("error at line(%d): %s\n", linenum, s);
  while (c != EOF && c != '\n') {
    c = getch();
  }
}

void getdef() {
  int c;
  int i;

  char def[MAXWORD];
  char dir[MAXWORD];
  char name[MAXWORD];

  skipblanks();

  if (!isalpha(getword(dir, MAXWORD))) {
    error(dir[0], __LINE__, "getdef: expecting a directive after #");
  } else if (strcmp(dir, "define") == 0) {
    skipblanks();

    if (!isalpha(getword(name, MAXWORD))) {
      error(name[0], __LINE__, "getdef: non-alpha - name expected");
    } else {
      skipblanks();

      for (i = 0; i < MAXWORD - 1; i++) {
        if ((def[i] = getch()) == EOF || def[i] == '\n') {
          break; /* end of definition */
        }
      }

      def[i] = '\0';

      if (i <= 0) { /* no definition */
        error('\n', __LINE__, "getdef: incomplete define");
      } else {
        install(name, def);
      }
    }
  } else if (strcmp(dir, "undef") == 0) {
    skipblanks();

    if (!isalpha(getword(name, MAXWORD))) {
      error(name[0], __LINE__, "getdef: non-alpha in undef");
    } else {
      undef(name);
    }
  } else {
    error(dir[0], __LINE__, "getdef: expecting a directive after #");
  }
}

/* skip blank and tab characters */
void skipblanks() {
  int c;
  while ((c = getch()) == ' ' || c == '\t') {
    ;
  }

  ungetch(c);
}

unsigned int hash(char *s) {
  unsigned int hashval;

  for (hashval = 0; *s != '\0'; s++) {
    hashval = *s + 31 * hashval;
  }

  return hashval % HASHSIZE;
}

struct nlist *lookup(char *s) {
  struct nlist *np;

  for (np = hashtable[hash(s)]; np != NULL; np = np->next) {
    if (strcmp(s, np->name) == 0) {
      return np;
    }
  }

  return NULL;
}

/* 将(name, defn) 加入到hashtable中 */
struct nlist *install(char *name, char *defn) {
  struct nlist *np;
  unsigned int hashval;

  if ((np = lookup(name)) == NULL) {
    np = (struct nlist *)malloc(sizeof(*np));
    if (np == NULL || (np->name = strdup(name)) == NULL) { /* TODO 什么意思? */
      return NULL;
    }

    hashval = hash(name);
    np->next = hashtable[hashval];
    hashtable[hashval] = np;
  } else {
    /* 释放前一个defn */
    free((void *)np->defn);
  }

  if ((np->defn = strdup(defn)) == NULL) {
    return NULL;
  }

  return np;
}

/* remove a name and definition from the table */
void undef(char *s) {
  unsigned int hashval;

  printf("undefine %s from hashtable\n", s);
  struct nlist *prev;
  struct nlist *np;

  prev = NULL;
  hashval = hash(s);

  for (np = hashtable[hashval]; np != NULL; np = np->next) {
    if (strcmp(s, np->name) == 0) {
      break;
    }

    prev = np; /* remember previous entry */
  }

  if (np != NULL) {     /* found name */
    if (prev == NULL) { /* first in the hash list */
      hashtable[hashval] = np->next;
    } else { /* elsewhere in the hash list */
      prev->next = np->next;
    }

    free((void *)np->name);
    free((void *)np->defn);
    free((void *)np); /* free allocated structure */
  }
}

void printtable() {
  for (int i = 0; i < HASHSIZE; i++) {
    for (struct nlist *current = hashtable[i]; current != NULL;
         current = current->next) {

      printf("%s\t\t%s\n", current->name, current->defn);
    }
  }

  printf("\n");
}

/*
page 127

以本节(6.6 表查找)介绍的函数为基础, 编写一个适合C语言程序使用的#define处理器的简单版本(即无参数的情况)
你会发现getch和ungetch函数非常有用

*/

#include <stdio.h>
#include <string.h>
#include <stdlib.h>
#include <ctype.h>

#include "../common-utils/getch-ungetch.c"

struct nlist {
  struct nlist *next; /* pointer to next node       */
  char *name;         /* defined name to substitute */
  char *defn;         /* real text                 */
};

#define HASHSIZE 101
#define MAXWORD  100

unsigned int hash(char *s);
struct nlist *lookup(char *s);

static struct nlist *hashtable[HASHSIZE];
struct nlist *install(char *name, char *defn);
void hashtable_dump();
void undef(char *s);
void getdef();
void error(int, char *);
int getword(char *, int);
void skipblanks();

/* simple version of #define processor */
int main(int argc, char const *argv[])
{
  // TODO how to test ?

  return 0;
}

int comment() {
  int c;
  while((c = getch()) != EOF) {
    if(c == '*') {
      if((c = getch()) == '/') {
        break;
      } else {
        ungetch(c);
      }
    }
  }

  return c;
}

/* get next word or character from input */
int getword(char *word, int limit) {
  int c;
  int d;
  char *w = word;

  while(isspace(c = getch())) {
    ;
  }

  if(c != EOF) {
    *w++ = c;
  }

  if(isalpha(c) || c == '_' | c == '#') {
    for( ; --limit > 0; w++) {
      if(!isalnum(*w = getch()) && *w != '_') {
        ungetch(*w);
        break;
      }
    }
  } else if(c == '\'' || c == '"') {
    for( ; --limit > 0; w++) {
      if((*w = getch()) == '\\') {
        *++w = getch();
      } else if(*w == c) {
        w++;
        break;
      } else if(*w == EOF) {
        break;
      }
    }
  } else if(c == '/') {
    if((d = getch()) == '#') {
      c = comment();
    } else {
      ungetch(d);
    }
  }

  *w = '\0';
  return c;
}

/* print error message and skip the rest of the line */
void error(int c, char *s) {
  printf("error: %s\n", s);
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
    error(dir[0], "getdef: expecting a directive after #");
  } else if (strcmp(dir, "define") == 0) {
    skipblanks();

    if (!isalpha(getword(name, MAXWORD))) {
      error(dir[0], "getdef: non-alpha - name expected");
    } else {
      skipblanks();

      for (i = 0; i < MAXWORD-1; i++) {
        if((def[i] = getch()) == EOF || def[i] == '\n') {
          break; /* end of definition */
        }
      }

      def[i] = '\0';

      if (i <= 0) { /* no definition */
        error('\n', "getdef: incomplete define");
      } else {
        install(name, def);
      }
    }
  } else if (strcmp(dir, "undef") == 0) {
    skipblanks();

    if(!isalpha(getword(name, MAXWORD))) {
      error(dir[0], "getdef: non-alpha in undef");
    } else {
      undef(name);
    }
  } else {
    error(dir[0], "getdef: expecting a directive after #");
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
    np = (struct nlist *) malloc(sizeof(*np));
    if (np == NULL || (np->name = strdup(name)) == NULL) { /* TODO 什么意思? */
      return NULL;
    }

    hashval = hash(name);
    np->next = hashtable[hashval];
    hashtable[hashval] = np;
  } else {
    /* 释放前一个defn */
    free((void *) np->defn);
  }

  if ((np->defn = strdup(defn)) == NULL) {
    return NULL;
  }

  return np;
}

void hashtable_dump() {
  for (int i = 0; i < HASHSIZE; i++) {
    for (struct nlist *current = hashtable[i]; current != NULL; current = current->next) {
      printf("%10s:%10s\n", current->name, current->defn);
    }
  }

  printf("\n");
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

  if (np != NULL) {    /* found name */
    if(prev == NULL) { /* first in the hash list */
      hashtable[hashval] = np->next;
    } else {           /* elsewhere in the hash list */
      prev->next = np->next;
    }

    free((void *) np->name);
    free((void *) np->defn);
    free((void *) np); /* free allocated structure */
  }
}


/*
page 编写函数undef, 它将从有lookup和install维护的表中删除一个变量名及其定义

ch6-Structures/text-substitution-demo01.c

*/

#include <stdio.h>
#include <string.h>
#include <stdlib.h>

struct nlist {
  struct nlist *next; /* pointer to next node       */
  char *name;         /* defined name to substitute */
  char *defn;         /* real text                 */
};

#define HASHSIZE 101

unsigned int hash(char *s);
struct nlist *lookup(char *s);

static struct nlist *hashtable[HASHSIZE];
struct nlist *install(char *name, char *defn);
void hashtable_dump();
void undef(char *s);

int main(int argc, char const *argv[])
{
  printf("before:\n=======================\n");
  install("date", "20220202");
  install("weather", "sunny");
  install("food", "noodles");
  hashtable_dump();

  undef("food");
  printf("after:\n=======================\n");
  hashtable_dump();

  return 0;
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


/*
page 编写函数undef, 它将从有lookup和install维护的表中删除一个变量名及其定义

ch6-Structures/text-substitution-demo01.c

*/

#include <stdio.h>
#include <string.h>
#include <stdlib.h>
#include <assert.h>

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
void undef(char *s);

int main(int argc, char const *argv[])
{
  assert(lookup("date") == NULL);
  assert(lookup("age") == NULL);
  assert(lookup("food") == NULL);

  install("date", "20220202");
  install("weather", "sunny");
  install("food", "noodles");

  assert(strcmp(lookup("date")->defn, "20220202") == 0);
  assert(strcmp(lookup("weather")->defn, "sunny") == 0);
  assert(strcmp(lookup("food")->defn, "noodles") == 0);

  undef("date");
  assert(lookup("date") == NULL);

  undef("weather");
  assert(lookup("weather") == NULL);

  undef("food");
  assert(lookup("food") == NULL);

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

/* remove a name and definition from the table */
void undef(char *s) {
  unsigned int hashval;

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


/*
page 125

编写一个表查找程序包的核心部分代码

这段代码很典型, 可以在宏处理器或者编译器的符号表管理例程中找到
例如: 遇到 `#define IN 1` 时, 需要把 `IN` 替换为 `1`

以下两个函数用来处理名字和替换文本
install(s, t)函数把名字s和替换文本t记录到某个表中, 其中s和t仅仅是字符串
lookup(s)函数在表中查找s, 若找到, 则返回指向该处的指针, 否则返回NULL

该算法采用散列查找方法

*/

#include <assert.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

/* linked list */
struct nlist {
  struct nlist *next; /* pointer to next node       */
  char *name;         /* defined name to substitute */
  char *defn;         /* real text                 */
};

#define HASHSIZE 101

unsigned int hash(char *s);
struct nlist *lookup(char *s);

static struct nlist *hashtable[HASHSIZE]; /* 指针表/指针的数组 */
struct nlist *install(char *name, char *defn);
void hashtable_dump();

int main(int argc, char const *argv[]) {
  assert(lookup("date") == NULL);
  assert(lookup("age") == NULL);

  install("date", "20220202");
  install("weather", "sunny");
  install("food", "noodles");

  assert(strcmp(lookup("date")->defn, "20220202") == 0);
  assert(strcmp(lookup("weather")->defn, "sunny") == 0);
  assert(strcmp(lookup("food")->defn, "noodles") == 0);

  hashtable_dump();
  return 0;
}

/* simple hash algorithm */
unsigned int hash(char *s) {
  unsigned int hashval;

  for (hashval = 0; *s != '\0'; s++) {
    hashval = *s + 31 * hashval;
  }

  return hashval % HASHSIZE;
}

struct nlist *lookup(char *s) {
  struct nlist *np;

  /* standard way of iterating through a linked list */
  /* for (ptr = head; ptr != NULL; ptr = ptr->next) */
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

void hashtable_dump() {
  printf("hashtable_dump:\n");
  printf("=====================\n");

  for (int i = 0; i < HASHSIZE; i++) {
    for (struct nlist *current = hashtable[i]; current != NULL;
         current = current->next) {
      printf("%10s:%10s\n", current->name, current->defn);
    }
  }

  printf("\n");
}

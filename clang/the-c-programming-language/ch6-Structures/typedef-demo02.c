/*
page 128

ch6-Structures/keyword-count-demo03.c

*/

#include <stdio.h>
#include <string.h>
#include <stdlib.h>

typedef struct tnode *Treeptr;

struct tnode {   /* tree node       */
  char *word;    /* pointer to word */
  int count;     /* word count      */
  Treeptr left;  /* left tree node  */
  Treeptr right; /* right tree node */
};

#define MAXWORD 100
Treeptr addtreex(Treeptr p, char *w);
void treeprint(Treeptr p);

// 统计单词出现的频率
int main(int argc, char const *argv[])
{
  Treeptr root;
  char word[MAXWORD];

  root = NULL;
  root = addtreex(root, "hello");
  root = addtreex(root, "world");
  root = addtreex(root, "world");
  root = addtreex(root, "world");
  root = addtreex(root, "hello");
  root = addtreex(root, "addtreex");
  root = addtreex(root, "addtreex");
  root = addtreex(root, "addtreex");
  root = addtreex(root, "addtreex");
  root = addtreex(root, "main");
  root = addtreex(root, "main");
  root = addtreex(root, "main");
  root = addtreex(root, "void");
  root = addtreex(root, "main");

  treeprint(root);

  return 0;
}

Treeptr talloc() {
  return (Treeptr ) malloc(sizeof(struct tnode));
}

Treeptr addtreex(Treeptr p, char *w) {
  int cond;

  if(p == NULL) {
    p = talloc();
    p->word = strdup(w);
    p->count = 1;
    p->left = NULL;
    p->right = NULL;
  } else if((cond = strcmp(w, p->word)) == 0) {
    p->count++;
  } else if(cond < 0) {
    p->left = addtreex(p->left, w);
  } else {
    p->right = addtreex(p->right, w);
  }

  return p;
}

void treeprint(Treeptr p) {
  if(p != NULL) {
    treeprint(p->left);
    printf("%4d %s\n", p->count, p->word);
    treeprint(p->right);
  }
}


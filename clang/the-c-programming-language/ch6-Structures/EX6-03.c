/*
page 125

编写一个交叉引用程序, 打印文档中所有单词的列表, 并且每个单词还有一个列表, 记录出现过该单词的行号
对the/and等非是一单词不予考虑

*/

#include <stdio.h>
#include <ctype.h>
#include <string.h>
#include <stdlib.h>

#define MAXWORD 100

/* linked list of line numbers */
struct linklist {
  int lnum;
  struct linklist *ptr;
};

struct tnode {
  char *word;             /* pointer to the text */
  struct linklist *lines; /* line numbers        */
  struct tnode *left;     /* left child          */
  struct tnode *right;    /* right child         */
};

struct tnode *addtreex(struct tnode *, char *, int);
int getword(char *, int);
int noiseword(char *);
void treeprint(struct tnode *);

/* cross referencer */
int main(int argc, char const *argv[])
{

  return 0;
}

struct tnode *talloc() {
  return (struct tnode *) malloc(sizeof(struct tnode));
}

struct linklist *lalloc() {
  return (struct linklist *) malloc(sizeof(struct linklist));
}


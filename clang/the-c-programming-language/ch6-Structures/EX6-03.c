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

// make a linked list node
struct linklist *lalloc() {
  return (struct linklist *) malloc(sizeof(struct linklist));
}

// add line number to the linked list
void addln(struct tnode *p, int linenum) {
  struct linklist *temp;

  temp = p->lines;
  while(temp->ptr != NULL && temp->lnum != linenum) {
    temp = temp->ptr;
  }

  if(temp->lnum != linenum) {
    temp->ptr = lalloc();
    temp->ptr->lnum = linenum;
    temp->ptr->ptr = NULL;
  }
}

void treeprint(struct tnode *p) {
  struct linklist *temp;

  if(p != NULL) {
    treeprint(p->left);
    printf("%10s: ", p->word);
    for(temp = p->lines; temp != NULL; temp = temp->ptr) {
      printf("%4d ", temp->lnum);
    }
    printf("\n");
    treeprint(p->right);
  }
}

// identify word as a noise word
int noiseword(char *w) {
  static char *nw[] = {
    "a",
    "an",
    "and",
    "are",
    "in",
    "is",
    "of",
    "or",
    "that",
    "the",
    "this",
    "to",
  };


  return 1;
}




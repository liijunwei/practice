/*
page 128

ch6-Structures/keyword-count-demo03.c

*/

#include <stdio.h>
#include <ctype.h>
#include <string.h>
#include <stdlib.h>

#include "../common-utils/getch-ungetch.c"

struct tnode {         /* tree node       */
  char *word;          /* pointer to word */
  int count;           /* word count      */
  struct tnode *left;  /* left tree node  */
  struct tnode *right; /* right tree node */
  /* 一个包含其自身实例的结构是非法的 */
  /* 但是上面的声明是合法的: left/right 指针的声明是指向tnode的指针, 而不是tnode实例本身 */
};

#define MAXWORD 100
struct tnode *addtreex(struct tnode *p, char *w);
void treeprint(struct tnode *p);
int getword(char *, int);

// bash ch6-Structures/keyword-count-test.sh

// 统计单词出现的频率
int main(int argc, char const *argv[])
{
  struct tnode *root;
  char word[MAXWORD];

  root = NULL;

  while(getword(word, MAXWORD) != EOF) {
    if(isalpha(word[0])) {
      root = addtreex(root, word);
    }
  }

  treeprint(root);

  return 0;
}

struct tnode *talloc() {
  return (struct tnode *) malloc(sizeof(struct tnode));
}

char *custom_strdup(char *s) {
  char *p;

  p = (char *) malloc(strlen(s) + 1);
  if(p != NULL) {
    strcpy(p, s);
  }

  return p;
}

struct tnode *addtreex(struct tnode *p, char *w) {
  int cond;

  if(p == NULL) {
    p = talloc();
    p->word = custom_strdup(w);
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

// 打印树p
void treeprint(struct tnode *p) {
  if(p != NULL) {
    treeprint(p->left);
    printf("%4d %s\n", p->count, p->word);
    treeprint(p->right);
  }
}



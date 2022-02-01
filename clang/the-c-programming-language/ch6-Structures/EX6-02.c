/*
page 125

编写一个程序, 用以读入一个C语言程序, 并按字母表顺序分组打印变量名, 要求每一组内各变量的前6个字符相同, 其余字符不同
字符串和注释的单词不予考虑
请将'6'作为一个可在命令行中设定的参数

*/

#include <stdio.h>
#include <ctype.h>
#include <string.h>
#include <stdlib.h>

#include "../common-utils/getch-ungetch.c"

struct tnode {
  char *word;
  int match;
  struct tnode *left;
  struct tnode *right;
};

#define MAXWORD 100
#define YES 1
#define NO 0

struct tnode *addtreex(struct tnode *, char *, int , int *);
void treeprint(struct tnode *);
int getword(char *, int);

/* print in alphabetic order each group of variable names */
/* identical in the first num characters (default 6) */
int main(int argc, char const *argv[])
{
  struct tnode *root;
  char word[MAXWORD];
  int found = NO;
  int num;

  num = (--argc && (*++argv)[0] == '-') ? atoi(argv[0] + 1) : 6;
  root = NULL;

  while(getword(word, MAXWORD) != EOF) {
    if(isalpha(word[0]) && strlen(word) >= num) {
      root = addtreex(root, word, num, &found);
    }

    found = NO;
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

// compare words and update p->match
int compare(char *s, struct tnode *p, int num, int *found) {
  int i;
  char *t = p->word;

  for(i = 0; *s == *t; i++, s++, t++) {
    if(*s == '\0') {
      return 0;
    }
  }

  if(i >= num) {
    *found = YES;
    p->match = YES;
  }

  return *s - *t;
}

struct tnode *addtreex(struct tnode *p, char *w, int num, int *found) {
  int cond;

  if(p == NULL) {
    p = talloc();
    p->word = custom_strdup(w);
    p->match = *found;
    p->left = NULL;
    p->right = NULL;
  } else if((cond = compare(w, p, num, found)) < 0) {
    p->left = addtreex(p->left, w, num, found);
  } else if(cond > 0) {
    p->right = addtreex(p->right, w, num, found);
  }

  return p;
}

int getword(char *word, int limit) {
  int c;
  char *w = word;

  while(isspace(c = getch())) {
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

void treeprint(struct tnode *p) {
  if(p != NULL) {
    treeprint(p->left);
    printf("%s\n", p->word);
    treeprint(p->right);
  }
}


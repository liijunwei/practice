/*
page 125

编写一个交叉引用程序, 打印文档中所有单词的列表, 并且每个单词还有一个列表,
记录出现过该单词的行号 对the/and等非是一单词不予考虑

*/

#include <ctype.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#include "../common-utils/getch-ungetch.c"

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
int main(int argc, char const *argv[]) {
  struct tnode *root;
  char word[MAXWORD];
  int linenum = 1;

  root = NULL;
  while (getword(word, MAXWORD) != EOF) {
    if (isalpha(word[0]) && noiseword(word) == -1) {
      root = addtreex(root, word, linenum);
    } else if (word[0] == '\n') {
      linenum++;
    }
  }

  treeprint(root);

  return 0;
}

int comment() {
  int c;
  while ((c = getch()) != EOF) {
    if (c == '*') {
      if ((c = getch()) == '/') {
        break;
      } else {
        ungetch(c);
      }
    }
  }

  return c;
}

int getword(char *word, int limit) {
  int c;
  int d;
  char *w = word;

  /* there you are!!!         */
  /* thanks "Gerrard Lindsay" */
  /* https://github.com/gerrard00/the-c-programming-language/blob/master/getword.c
   */
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
  } else if (c == '/') {
    if ((d = getch()) == '#') {
      c = comment();
    } else {
      ungetch(d);
    }
  }

  *w = '\0';
  return c;
}

struct tnode *talloc() {
  return (struct tnode *)malloc(sizeof(struct tnode));
}

// make a linked list node
struct linklist *lalloc() {
  return (struct linklist *)malloc(sizeof(struct linklist));
}

// add line number to the linked list
void addln(struct tnode *p, int linenum) {
  struct linklist *temp;

  temp = p->lines;
  while (temp->ptr != NULL && temp->lnum != linenum) {
    temp = temp->ptr;
  }

  if (temp->lnum != linenum) {
    temp->ptr = lalloc();
    temp->ptr->lnum = linenum;
    temp->ptr->ptr = NULL;
  }
}

void treeprint(struct tnode *p) {
  struct linklist *temp;

  if (p != NULL) {
    treeprint(p->left);
    printf("%10s: ", p->word);
    for (temp = p->lines; temp != NULL; temp = temp->ptr) {
      printf("%4d ", temp->lnum);
    }
    printf("\n");
    treeprint(p->right);
  }
}

// identify word as a noise word
int noiseword(char *w) {
  static char *nw[] = {
      "a",  "an", "and",  "are", "in",   "is",
      "of", "or", "that", "the", "this", "to",
  };

  int cond;
  int mid;
  int low = 0;
  int high = sizeof(nw) / sizeof(char *) - 1;

  while (low <= high) {
    mid = (low + high) / 2;
    if ((cond = strcmp(w, nw[mid])) < 0) {
      high = mid - 1;
    } else if (cond > 0) {
      low = mid + 1;
    } else {
      return mid;
    }
  }

  return -1;
}

// add a node with w, at or below p
struct tnode *addtreex(struct tnode *p, char *w, int linenum) {
  int cond;

  if (p == NULL) {
    p = talloc();
    p->word = strdup(w);
    p->lines = lalloc();
    p->lines->lnum = linenum;
    p->lines->ptr = NULL;
    p->left = NULL;
    p->right = NULL;
  } else if ((cond = strcmp(w, p->word)) == 0) {
    addln(p, linenum);
  } else if (cond < 0) {
    p->left = addtreex(p->left, w, linenum);
  } else {
    p->right = addtreex(p->right, w, linenum);
  }

  return p;
}

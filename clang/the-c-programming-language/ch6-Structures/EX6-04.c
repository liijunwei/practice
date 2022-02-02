/*
page 125

编写一个程序, 根据单词出现的频率, 按降序打印输入的各个不同单词, 并在每个单词的前面标上它的出现次数


*/

#include <stdio.h>
#include <ctype.h>
#include <string.h>
#include <stdlib.h>

#include "../common-utils/getch-ungetch.c"

#define MAXWORD  100
#define NDISTNCT 1000

struct tnode {
  char *word; /* pointer to the text   */
  int count;  /* number of occurrences */
  struct tnode *left;
  struct tnode *right;
};

struct tnode *addtree(struct tnode *, char *);
int getword(char *, int);
void sortlist();
void treestore(struct tnode *);

struct tnode *list[NDISTNCT]; /* pointer to tree nodes */
int ntn = 0;                  /* number of tree nodes  */

/* print distinct words sorted in decreasing order of frequency */
int main(int argc, char const *argv[])
{

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

int getword(char *word, int limit) {
  int c;
  int d;
  char *w = word;

  while(isspace(c = getch()) && c != '\n') {
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

/* store in list[] pointers to tree nodes */
void treestore(struct tnode *p) {
  if (p != NULL) {
    treestore(p->left);
    if (ntn < NDISTNCT) {
      list[ntn] = p;
      ntn++;
    }
    treestore(p->right);
  }
}

/* sort list of pointers to tree nodes */
void sortlist() {
  int gap;
  int i;
  int j;

  struct tnode *temp;

  for (gap = ntn/2; gap > 0; gap /= 2) {
    for (i = gap; i < ntn; i++) {
      for (j = i - gap; j >= 0; j -= gap) {
        if ((list[j]->count) >= (list[j+gap]->count)) {
          break;
        }

        temp = list[j];
        list[j] = list[j+gap];
        list[j+gap] = temp;
      }
    }
  }
}




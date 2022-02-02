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

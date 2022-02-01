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

struct tnode {
  char *word;
  int match;
  struct tnode *left;
  struct tnode *right;
};

#define MAXWORD 100
#define YES 1
#define NO 0

struct tnode *address(struct tnode *, char *, int , int *);
void treeprint(struct tnode *);
int getword(char *, int);

int main(int argc, char const *argv[])
{

  return 0;
}

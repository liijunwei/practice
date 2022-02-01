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

int main(int argc, char const *argv[])
{

  return 0;
}

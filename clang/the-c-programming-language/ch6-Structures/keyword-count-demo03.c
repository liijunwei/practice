/*
page 121

处理更一般的问题: 统计输入中所有单词的出现频次

使用二叉树存储数据: 任何节点最多有两个子树, 也可能只有一个子树或0个子树
对节点的所有操作要保证, 任何节点的左子树只包含按字典顺序小于该节点的那些单词
右子树只包含按字典顺序大于该节点中单词的那些单词

now is the time for all good men to come the aid of their party
                          now
                    is          the
                for    men   of          time
            all    good        party their   to
        aid   come
*/

#include <stdio.h>
#include <ctype.h>
#include <string.h>

#include "../common-utils/getch-ungetch.c"

struct tnode {         /* tree node       */
  char *word;          /* pointer to word */
  int count;           /* word count      */
  struct tnode *left;  /* left tree node  */
  struct tnode *right; /* right tree node */
};

int main(int argc, char const *argv[])
{


  return 0;
}
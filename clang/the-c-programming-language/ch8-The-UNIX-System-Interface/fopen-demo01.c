/*
page 155

man 3 fopen
*/

#define NULL     0
#define EOF      (-1)
#define BUFSIZ   1024
#define OPEN_MAX 20 /* 一次最多可打开的文件数 */

typedef struct _iobuf { /* 只供标准库中其他函数使用的名字以下划线开始 */
  int cnt;    /* 剩余的字符数     */
  char *ptr;  /* 下一个字符的位置 */
  char *base; /* 缓冲区的位置     */
  int flag;   /* 文件访问模式     */
  int fd;     /* 文件描述符       */
} FILE;

int main(int argc, char const *argv[]) {


  return 0;
}

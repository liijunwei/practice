/*
page 157

标准库函数 `int fseek(FILE *fp, long offset, int origin)` 类似于函数 `lseek`, 所不同的是, 该函数中的fp
是一个文件指针而不是文件描述符, 且返回值是一个int类型的状态而非位置值.

编写函数fseek, 并确保该函数与库中其他函数使用的缓冲能够协同工作

ch8-The-UNIX-System-Interface/fopen-demo01.c

*/

#include <unistd.h>
#include <fcntl.h>
#include <stdlib.h>

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

extern FILE _iob[OPEN_MAX];

#define stdin  (&_iob[0])
#define stdout (&_iob[1])
#define stderr (&_iob[2])

enum _flags {
  _READ  = 01,  /* 以读的方式打开文件 */
  _WRITE = 02,  /* 以写的方式打开文件 */
  _UNBUF = 04,  /* 不对文件进行缓冲   */
  _EOF   = 010, /* 已到文件的末尾     */
  _ERR   = 020, /* 该文件发生错误     */
};
int custom_fseek(FILE *fp, long offset, int origin) {
  unsigned int nc;
  long rc = 0;

  if (fp->flag & _READ) {
    if (origin == 1) {
      offset -= fp->cnt;
    }

    rc = lseek(fp->fd, offset, origin);
    fp->cnt = 0;
  } else if (fp->flag & _WRITE) {
    if ((nc = fp->ptr - fp->base) > 0) {
      if (write(fp->fd, fp->base, nc) != nc) {
        rc = -1;
      }
    }

    if (rc != -1) {
      rc = lseek(fp->fd, offset, origin);
    }
  }

  return (rc == 1) ? -1 : 0;
}

int main(int argc, char const *argv[])
{

  return 0;
}

/*
page 155

man 3 fopen
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

int _fillbuf(FILE *fp);
int _flushbuf(int a, FILE *fp);

#define feof(p)   (((p)->flag & _EOF) != 0)
#define ferror(p) (((p)->flag & _ERR) != 0)
#define fileno(p) ((p)->fd)

#define getc(p) (--(p)->cnt >= 0 \
              ? (unsigned char) *(p)->ptr++ : _fillbuf(p))

#define putc(x, p) (--(p)->cnt >= 0 \
              ? *(p)->ptr++ = (x) : _flushbuf((x), p))

#define getchar() getc(stdin)
#define putchar(x) putc((x), stdout)

#define PERMS 0666

FILE *custom_fopen(char *name, char *mode) {
  int fd;
  FILE *fp;

  if (*mode != 'r' && *mode != 'w' && *mode != 'a') {
    return NULL;
  }

  for (fp = _iob; fp < _iob + OPEN_MAX; fp++) {
    if ((fp->flag & (_READ | _WRITE)) == 0) {
      break;
    }
  }

  if (fp >= _iob + OPEN_MAX) {
    return NULL;
  }

  if (*mode == 'w') {
    fd = creat(name, PERMS);
  } else if (*mode == 'a') {
    if((fd = open(name, O_WRONLY, 0)) == -1) {
      fd = creat(name, PERMS);
    }

    lseek(fd, 0L, 2);
  } else {
    fd = open(name, O_RDONLY, 0);
  }

  if (fd == -1) {
    return NULL;
  }

  fp->fd = fd;
  fp->cnt = 0;
  fp->base = NULL;
  fp->flag = (*mode == 'r') ? _READ : _WRITE;

  return fp;
}

int _fillbuf(FILE *fp) {
  int bufsize;

  if ((fp->flag & (_READ | _EOF | _ERR)) != _READ) {
    return EOF;
  }

  bufsize = (fp->flag & _UNBUF) ? 1 : BUFSIZ;

  if (fp->base == NULL) {
    if ((fp->base = (char *) malloc(bufsize)) == NULL) {
      return EOF;
    }
  }

  fp->ptr = fp->base;
  fp->cnt = read(fp->fd, fp->ptr, bufsize);

  if (--fp->cnt < 0) {
    if (fp->cnt == -1) {
      fp->flag |= _EOF;
    } else {
      fp->flag |= _ERR;
    }

    fp->cnt = 0;

    return EOF;
  }

  return (unsigned char) *fp->ptr++;
}

int _flushbuf(int a, FILE *fp) {

  return 0;
}

FILE _iob[OPEN_MAX] = {
  {0, (char *) 0, (char *) 0, _READ,             0}, /* stdin  */
  {0, (char *) 0, (char *) 0, _WRITE,            1}, /* stdout */
  {0, (char *) 0, (char *) 0, (_WRITE | _UNBUF), 2}, /* stderr */
};

int main(int argc, char const *argv[]) {
  int c;

  return 0;
}

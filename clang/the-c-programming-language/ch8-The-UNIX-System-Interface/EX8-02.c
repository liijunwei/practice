/*
page 157

ch8-The-UNIX-System-Interface/fopen-demo01.c
用字段代替显式的按位操作, 重写fopen和_fillbuf函数; 比较相应代码的程度和执行速度

TODO unclear
*/

#include <unistd.h>
#include <fcntl.h>
#include <stdlib.h>

#define EOF      (-1)
#define BUFSIZ   1024
#define OPEN_MAX 20 /* 一次最多可打开的文件数 */

struct flag_field {
  unsigned int is_read:  1;
  unsigned int is_write: 1;
  unsigned int is_unbuf: 1;
  unsigned int is_buf:   1;
  unsigned int is_eof:   1;
  unsigned int is_err:   1;
};

typedef struct _iobuf { /* 只供标准库中其他函数使用的名字以下划线开始 */
  int cnt;    /* 剩余的字符数     */
  char *ptr;  /* 下一个字符的位置 */
  char *base; /* 缓冲区的位置     */
  struct flag_field flag;   /* 文件访问模式     */
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

/* open file, return file ptr */
FILE *custom_fopen(char *name, char *mode) {
  int fd;
  FILE *fp;

  if (*mode != 'r' && *mode != 'w' && *mode != 'a') {
    return NULL;
  }

  for (fp = _iob; fp < _iob + OPEN_MAX; fp++) {
    if (fp->flag.is_read == 0 && fp->flag.is_write == 0) {
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
  fp->flag.is_unbuf = 0;
  fp->flag.is_buf = 1;
  fp->flag.is_eof = 0;
  fp->flag.is_err = 0;

  if (*mode == 'r') {
    fp->flag.is_read = 1;
    fp->flag.is_write = 0;
  } else {
    fp->flag.is_read = 0;
    fp->flag.is_write = 1;
  }

  return fp;
}

int _fillbuf(FILE *fp) {
  int bufsize;

  if (fp->flag.is_read == 0 || fp->flag.is_eof == 1 || fp->flag.is_err == 1) {
    return EOF;
  }

  bufsize = (fp->flag.is_unbuf == 1) ? 1 : BUFSIZ;

  if (fp->base == NULL) {
    if ((fp->base = (char *) malloc(bufsize)) == NULL) {
      return EOF;
    }
  }

  fp->ptr = fp->base;
  fp->cnt = read(fp->fd, fp->ptr, bufsize);

  if (--fp->cnt < 0) {
    if (fp->cnt == -1) {
      fp->flag.is_eof = 1;
    } else {
      fp->flag.is_err = 1;
    }

    fp->cnt = 0;

    return EOF;
  }

  return (unsigned char) *fp->ptr++;
}

/* EX8-03 */
int _flushbuf(int x, FILE *fp) {
  unsigned int nc;
  int bufsize;

  if (fp < _iob || fp >= _iob + OPEN_MAX) {
    return EOF;
  }

  if (fp->flag.is_eof == 1) {
    return EOF;
  }

  bufsize = (fp->flag.is_unbuf) ? 1 : BUFSIZ;

  if (fp->base == NULL) {
    if ((fp->base = (char *) malloc(bufsize)) == NULL) {
      fp->flag.is_err = 1;
      return EOF;
    }
  } else {
    nc = fp->ptr - fp->base;
    if (write(fp->fd, fp->base, nc) != nc) {
      fp->flag.is_err = 1;
      return EOF;
    }
  }

  fp->ptr = fp->base;
  *fp->ptr++ = (char) x;
  fp->cnt = bufsize - 1;

  return x;
}

FILE _iob[OPEN_MAX] = {
  {0, (char *) 0, (char *) 0, {1, 0, 0, 0, 0, 0}, 0}, /* stdin  */
  {0, (char *) 0, (char *) 0, {0, 1, 0, 0, 0, 0}, 1}, /* stdout */
  {0, (char *) 0, (char *) 0, {0, 1, 1, 0, 0, 0}, 2}, /* stderr */
};

/* EX8-03 */
int fflush(FILE *fp) {
  int rc = 0;

  if (fp < _iob || fp >= _iob + OPEN_MAX) {
    return EOF;
  }

  if (fp->flag.is_write == 1) {
    rc = _flushbuf(0, fp);
  }

  fp->ptr = fp->base;
  fp->cnt = (fp->flag.is_unbuf == 1) ? 1 : BUFSIZ;

  return rc;
}

/* EX8-03 */
int fclose(FILE *fp) {
  int rc;

  if ((rc = fflush(fp)) != EOF) {
    free(fp->base);
    fp->ptr = NULL;
    fp->cnt = 0;
    fp->base = NULL;
    fp->flag.is_read = 1;
    fp->flag.is_write = 1;
  }

  return rc;
}

// run ch8-The-UNIX-System-Interface/EX8-02.c
int main(int argc, char const *argv[]) {
  FILE *fp;
  fp = custom_fopen("./ch8-The-UNIX-System-Interface/EX8-02.c", "r");

  int c;

  while ((c = getc(fp)) != EOF) {
    putchar(c);
  }

  fclose(fp);

  return 0;
}

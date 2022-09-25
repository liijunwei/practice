#include "csapp.h"

/* 子进程汇集成父进程的描述符表, 以及所有进程共享的同一个打开文件表 */
/* 因此 描述符fd在父子进程中都指向同一个打开的文件表表项 */
/* 当子进程读取文件的第一个字节时, 文件位置加一, 因此子进程会读取第二个字节 */
int main(int argc, char const *argv[]) {
  int fd;
  char c;

  fd = Open("foobar.txt", O_RDONLY, 1);

  if (Fork() == 0) {
    Read(fd, &c, 1);
    exit(0);
  }

  Wait(NULL);
  Read(fd, &c, 1);

  printf("c = %c\n", c);

  exit(0);
}

/*
page 154

man 2 lseek
*/

#include <stdio.h>
#include <unistd.h>
#include <fcntl.h>

/* 从pos位置处读入n个字节 */
int get(int fd, long pos, char *buf, int n) {
  if (lseek(fd, pos, 0) >= 0) {
    return read(fd, buf, n); /* 移动到pos处 */
  } else {
    return -1;
  }
}

int main(int argc, char const *argv[]) {
  int fd = open("./ch8-The-UNIX-System-Interface/lseek-demo01.c", 0, O_RDONLY);
  char buf[BUFSIZ];

  printf("%d bytes data read\n", get(fd, 0L, buf, 40));
  printf("%s\n\n", buf);

  printf("%d bytes data read\n", get(fd, 0L, buf, 400));
  printf("%s\n\n", buf);

  printf("%d bytes data read\n", get(fd, 0L, buf, 4000));
  printf("%s\n\n", buf);

  return 0;
}

// 9.5 编写一个从程序, 使用mmap将一个任意大小的磁盘文件复制到stdout. 输入文件的名字必须作为一个命令行参数来传递

#include "csapp.h"

void mmapcopy(int fd, int size) {
  char *bufp; // ptr to memory-mapped VM area

  bufp = Mmap(NULL, size, PROT_READ, MAP_PRIVATE, fd, 0);
  Write(1, bufp, size);
  return;
}

// mmapcopy driver
int main(int argc, char const *argv[]) {
  struct stat stat;
  int fd;

  // check for required command-line argument
  if(argc != 2) {
    printf("Usage: %s <filename>\n", argv[0]);
    exit(0);
  }

  // copy the input argument to stdout
  fd = Open(argv[1], O_RDONLY, 0);
  fstat(fd, &stat);
  mmapcopy(fd, stat.st_size);

  exit(0);
}

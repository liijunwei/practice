// 9.5 编写一个从程序, 使用mmap将一个任意大小的磁盘文件复制到stdout. 输入文件的名字必须作为一个命令行参数来传递

#include "csapp.h"

void mmapcopy(int fd, int size) {
  char *bufp; // ptr to memory-mapped VM area

  bufp = Mmap(NULL, size, PROT_READ, MAP_PRIVATE, fd, 0);
  Write(1, bufp, size);
  return;
}

int main(int argc, char const *argv[]) {


  return 0;
}

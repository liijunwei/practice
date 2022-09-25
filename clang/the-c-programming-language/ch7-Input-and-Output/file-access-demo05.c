/*
page 144

char *fgets(char *line, int maxline, FILE *fp)
    fgets函数从fp指向的文件中读取下一个输入行(包括换行符),
并将他们存在字符数组line里, 最多可读取maxline-1个字符,
读取到的行将以'\0'结尾保存到数组里 通常情况下, fgets返回line,
但如果遇到了文件结尾或者发生了错误, 则返回NULL

int fputs(char *line, FILE *fp)
    fputs函数将一个字符串(不需要包含换行符)写入到一个文件中
    如果发生错误, 该函数将返回EOF, 否则返回一个非负值

*/

#include <stdio.h>

int main(int argc, char const *argv[]) {
  FILE *fp;
  fp = fopen("./ch7-Input-and-Output/file-access-demo05.c", "r");

  char line[100];

  while (fgets(line, 100, fp) != NULL) {
    printf("%s", line);
  }

  fclose(fp);

  return 0;
}

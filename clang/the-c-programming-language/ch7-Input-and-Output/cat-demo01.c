/*
page 142

int fscanf(FILE *fp, char *fmt, ...)
int fprintf(FILE *fp, char *fmt, ...)

编写将多个文件连接起来的cat程序
该程序的设计思路和其他许多程序类似:
  如果有命令行参数, 参数将被解释为文件名, 并按顺序逐个处理
  如果没有参数, 则处理标准输入
*/

#include <stdio.h>

void filecopy(FILE *ifp, FILE *ofp);

// gcc ch7-Input-and-Output/cat-demo01.c && ./a.out 0.goal.md README.md demo.c
/* 连接多个文件(V1) */
int main(int argc, char const *argv[]) {
  FILE *fp;

  if (argc == 1) {
    filecopy(stdin, stdout);
  } else {
    while (--argc > 0) {
      if ((fp = fopen(*++argv, "r")) == NULL) {
        printf("cat: can't open %s\n", *argv);
        return 1;
      } else {
        filecopy(fp, stdout);
        fclose(fp);
      }
    }
  }

  return 0;
}

void filecopy(FILE *ifp, FILE *ofp) {
  int c;

  while ((c = getc(ifp)) != EOF) {
    putc(c, ofp);
  }
}

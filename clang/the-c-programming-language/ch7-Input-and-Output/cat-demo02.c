/*
page 143

cat程序的功能不完善, 问题在于 如果因某种原因导致其中的一个文件无法访问,
相应的诊断信息要在该连接输出的末尾才能打印出来 当输出到屏幕时,
这种处理方法尚可接受; 但如果输出到一个文件或者通过管道输出到另一个程序时,
就无法接受了(算是脏数据)

为了更好地处理这种情况, 另一个输出流以与stdin和stdout相同的方式分派给程序,
即stderr; 即使对标准输出进行了重定向, 写到stderr中的输出通常也会显示在屏幕上

下面修改cat程序, 将其出错信息写到标准错误文件上

Q: exit(0) 和 return 0 有什么区别?
A: 在主函数main里, return expr <=> exit(expr); 但是使用函数exit有一个优点,
他可以从其他函数中调用 foobar/foo04.c
*/

#include <stdio.h>
#include <stdlib.h>
#include <time.h>

void filecopy(FILE *ifp, FILE *ofp);

// gcc ch7-Input-and-Output/cat-demo02.c && ./a.out demo.c # OK
// gcc ch7-Input-and-Output/cat-demo02.c && ./a.out demo.ccc # ERROR
/* 连接多个文件(V2) */
int main(int argc, char const *argv[]) {
  clock_t start_time = clock();

  FILE *fp;
  char const *progname = argv[0];

  if (argc == 1) {
    filecopy(stdin, stdout);
  } else {
    while (--argc > 0) {
      if ((fp = fopen(*++argv, "r")) == NULL) {
        fprintf(stderr, "%s: can't open %s\n", progname, *argv);
        exit(1);
      } else {
        filecopy(fp, stdout);
        fclose(fp);
      }
    }
  }

  if (ferror(stdout)) {
    fprintf(stderr, "%s: error writing stdout\n", progname);
  }

  double elapsed_time = (double)(clock() - start_time) / CLOCKS_PER_SEC;
  printf("Done in %f seconds\n", elapsed_time);

  exit(0);
}

void filecopy(FILE *ifp, FILE *ofp) {
  int c;

  while ((c = getc(ifp)) != EOF) {
    putc(c, ofp);
  }
}

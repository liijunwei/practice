/*
page 145

编写一个程序, 以打印一个文件集合, 每个文件从新的一页开始打印,
并且打印每个文件相应的标题和页数

TODO not clear
*/

#include <stdio.h>
#include <stdlib.h>

#define MAXBOT 3    /* maximum number of lines at bottom page */
#define MAXHDR 5    /* maximum number of lines at head page */
#define MAXLINE 300 /* maximum size of one line */
#define MAXPAGE 50  /* maximum number of lines one one page */

void fileprint(FILE *fp, char const *fname);
int heading(char const *fname, int pageno);

/* print files: each new one on a new page */
int main(int argc, char const *argv[]) {
  FILE *fp;

  if (argc == 1) {
    fileprint(stdin, " ");
  } else {
    while (--argc > 0) {
      if ((fp = fopen(*++argv, "r")) == NULL) {
        fprintf(stderr, "print: can't open %s\n", *argv);
        exit(1);
      } else {
        fileprint(fp, *argv);
        fclose(fp);
      }
    }
  }

  return 0;
}

/* print file fname */
void fileprint(FILE *fp, char const *fname) {
  int lineno;
  int pageno = 1;

  char line[MAXLINE];
  lineno = heading(fname, pageno++);

  while (fgets(line, MAXLINE, fp) != NULL) {
    if (lineno == 1) {
      fprintf(stdout, "\f");
      lineno = heading(fname, pageno++);
    }

    fputs(line, stdout);

    if (++lineno > MAXPAGE - MAXBOT) {
      lineno = 1;
    }
  }

  fprintf(stdout, "\f"); /* page eject after the file */
}

/* put heading end enough blank lines */
int heading(char const *fname, int pageno) {
  int line_num = 3;

  fprintf(stdout, "\n\n");
  fprintf(stdout, "%s page %d\n", fname, pageno);

  while (line_num++ < MAXHDR) {
    fprintf(stdout, "\n");
  }

  return line_num;
}

#include <stdio.h>

/*
page 102

ch1-A-Tutorial-Introduction/EX1-20.c
ch1-A-Tutorial-Introduction/EX1-21.c

修改程序entab和detab(第一章练习中编写的函数)的功能, 使他们接收一组作为参数的制表符停止位
如果启动程序时不带参数, 则使用默认的制表符停止位设置

TODO 过...

*/

#define MAXLINE  100
#define TABING   8
#define YES      1
#define NO       0
#define CHAR_TAB '*' // '\t'

void settab(int argc, char const *argv[], char *tab);
void entab(char *tab);
int tabpos(int pos, char *tab);

// replace strings of blanks with tabs
int main(int argc, char const *argv[])
{
  char tab[MAXLINE + 1];

  settab(argc, argv, tab); // initialize tab stops
  entab(tab);              // replace blanks with tab

  return 0;
}

void settab(int argc, char const *argv[], char *tab){

}

void entab(char *tab){

}

int tabpos(int pos, char *tab){

  return 0;
}

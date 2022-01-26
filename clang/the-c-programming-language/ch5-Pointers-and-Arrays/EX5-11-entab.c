#include <stdio.h>
#include <stdlib.h>

/*
page 102

ch1-A-Tutorial-Introduction/EX1-20.c
ch1-A-Tutorial-Introduction/EX1-21.c

修改程序entab和detab(第一章练习中编写的函数)的功能, 使他们接收一组作为参数的制表符停止位
如果启动程序时不带参数, 则使用默认的制表符停止位设置

TODO 过...

*/

#define MAXLINE    100
#define TABING     8
#define YES        1
#define NO         0
#define CHAR_BLANK '_' // ' '
#define CHAR_TAB   '*' // '\t'

void settab(int argc, char const *argv[], char *tab);
void entab(char *tab);
void detab(char *tab);
int tabpos(int pos, char *tab);

// ./a.out 1 2 3 4 7

// replace strings of blanks with tabs
int main(int argc, char const *argv[])
{
  char tab[MAXLINE + 1];

  settab(argc, argv, tab); // initialize tab stops
  entab(tab);              // replace blanks with tab

  return 0;
}

// set tab stops in array tab
void settab(int argc, char const *argv[], char *tab){
  int i;
  int pos;

  if(argc <= 1){ // default tab stops
    for(i = 1; i <= MAXLINE; i++){
      if(i % TABING == 0){
        tab[i] = YES;
      } else {
        tab[i] = NO;
      }
    }
  } else { // user provided tab stops
    for(i = 1; i <= MAXLINE; i++){
      tab[i] = NO; // turn off all tab stops
    }

    while(--argc > 0){ // walk through argument list
      pos = atoi(*++argv);
      if(pos > 0 && pos <= MAXLINE){
        tab[pos] = YES;
      }
    }
  }
}

void entab(char *tab){
  int c;
  int pos;
  int nb = 0;
  int nt = 0;

  for(pos = 1; (c = getchar()) != EOF; pos++){
    if(c == ' '){
      if(tabpos(pos, tab) == NO){
        nb++;
      } else {
        nb = 0;
        nt++;
      }
    } else {
      for(; nt > 0; nt--){
        putchar(CHAR_TAB);
      }

      if(c == '\t'){
        nb = 0;
      } else {
        for(; nb > 0; nb --){
          putchar(CHAR_BLANK);
        }
      }

      putchar(c);

      if(c == '\n'){
        pos = 0;
      } else if (c == '\t'){
        while(tabpos(pos, tab) != YES){
          pos++;
        }
      }
    }
  }
}

// determine if pos is at a tab stop
int tabpos(int pos, char *tab){
  return pos > MAXLINE ? YES : tab[pos];
}

void detab(char *tab){
  int c;
  int pos;
  while((c = getchar()) != EOF){
    if(c == '\t'){
      do
        putchar(' ');
      while(tabpos(pos++, tab) != YES);
    } else if (c == '\n') {
      putchar(c);
      pos = 1;
    } else {
      putchar(c);
      pos++;
    }
  }
}


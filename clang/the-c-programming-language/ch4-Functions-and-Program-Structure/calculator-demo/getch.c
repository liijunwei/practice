#include <stdio.h>
#include "./calc.h"

#define BUFSIZE 100

char buf[BUFSIZE];
int bufp = 0; // buf中下一空闲的位置

// 取一个字符(可能是压回的字符)
int getch(void){
  return (bufp > 0) ? buf[--bufp] : getchar();
}

// 把字符压回输入中
void ungetch(int c){
  if(bufp >= BUFSIZE) {
    printf("Ungetch, too many characters\n");
  } else {
    buf[bufp++] = c;
  }
}

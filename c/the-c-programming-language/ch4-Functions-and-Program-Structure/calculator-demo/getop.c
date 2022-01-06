#include <stdio.h>
#include <ctype.h>
#include "./calc.h"

int getop(char s[]){
  int i;
  int c;

  while((s[0] = c = getch()) == ' ' || c == '\t'){
    ;
  }

  s[1] = '\0';
  if(!isdigit(c) && c != '.') {
    return c; // 不是数
  }

  i = 0;
  if(isdigit(c)) {
    // 整数部分
    while(isdigit(s[++i] = c = getch())){
      ;
    }
  }

  if(c == '.') {
    // 小数部分
    while(isdigit(s[++i] = c = getch())){
      ;
    }
  }

  s[i] = '\0';

  if(c != EOF) {
    ungetch(c);
  }

  return NUMBER;
}

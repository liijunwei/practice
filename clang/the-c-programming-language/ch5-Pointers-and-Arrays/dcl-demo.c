/*
page 107

根据声明符的语法对声明进行分析

*/

#include <stdio.h>
#include <string.h>
#include <ctype.h>

#include "../common-utils/getch-ungetch.c"

#define MAXTOKEN 100

enum {
  NAME,
  PARENS,
  BRACKETS
};

void dcl();
void dirdcl();
int gettoken();

int tokentype;
char token[MAXTOKEN];
char name[MAXTOKEN];
char datatype[MAXTOKEN];
char out[1000];

/* 将声明转为文字描述 */
int main(int argc, char const *argv[])
{
  while(gettoken() != EOF) {
    strcpy(datatype, token);
    out[0] = '\0';
    dcl();

    if(tokentype != '\n') {
      printf("Syntax error\n");
    }

    printf("%s %s %s\n", name, out, datatype);
  }

  return 0;
}

void dcl() {
  int ns;
  for(ns = 0; gettoken() == '*';) {
    ns++;
  }

  dirdcl();

  while(ns-- > 0) {
    strcat(out, " pointer to");
  }
}

void dirdcl() {

}

int gettoken() {
  int c;
  char *p = token;

  while((c = getch()) == ' ' || c == '\t') {
    ;
  }

  if(c == '(') {
    if((c = getch()) == ')') {
      strcpy(token, "()");
      return tokentype = PARENS;
    } else {
      ungetch(c);
      return tokentype = '(';
    }
  } else if(c == '[') {
    for(*p++ = c; (*p++ = getch()) != ']';) {
      ;
    }

    *p = '\0';
    return tokentype = BRACKETS;
  } else if(isalpha(c)) {
    for(*p++ = c; isalnum(c = getch());) {
      *p++ = c;
    }
    *p = '\0';
    ungetch(c);
    return tokentype = NAME;
  } else {
    return tokentype = c;
  }

  return 1;
}


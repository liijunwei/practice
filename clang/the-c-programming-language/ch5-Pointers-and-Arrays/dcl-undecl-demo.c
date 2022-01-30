/*
page 109

根据声明符的语法对声明进行分析

TODO 补充注释
*/

#include <stdio.h>

#include "../common-utils/getch-ungetch.c"

#define MAXTOKEN 100

enum {
  NAME,
  PARENS,
  BRACKETS
};

int gettoken();

int tokentype;
char token[MAXTOKEN];
char out[1000];

// x is a function returning a pointer to an array of pointers to functions returning char
// <=>
// f () * [] * () char
// test:
// ch5-Pointers-and-Arrays/dcl-undecl-demo-test.sh

/* undecl 函数, 将文字描述转换为声明 */
int main(int argc, char const *argv[])
{
  int type;
  char temp[MAXTOKEN];

  while(gettoken() != EOF) {
    strcat(out, token);

    while((type = gettoken()) != '\n') {
      if(type == PARENS || type == BRACKETS) {
        strcat(out, token);
      } else if(type == '*') {
        sprintf(temp, "(*%s)", out);
        strcpy(out, temp);
      } else if(type == NAME) {
        sprintf(temp, "%s %s", token, out);
        strcpy(out, temp);
      } else {
        printf("invalid input at %s\n", token);
      }
    }

    printf("%s\n", out);
  }

  return 0;
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
}


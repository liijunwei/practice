/*
page 109

修改undcl程序, 使他把文字描述转换为声明的过程中不会生成多余的圆括号

*/

#include <stdio.h>

#include "../common-utils/getch-ungetch.c"

#define MAXTOKEN 100

enum { NAME, PARENS, BRACKETS };

enum { NO, YES };

int gettoken();
int nexttoken();

int tokentype;
char token[MAXTOKEN];
char out[1000];
int prevtoken;

// bash ch5-Pointers-and-Arrays/EX5-19-test.sh

/* undecl: convert word description to declaration */
int main(int argc, char const *argv[]) {
  int type;
  char temp[MAXTOKEN];

  while (gettoken() != EOF) {
    strcpy(out, token);

    while ((type = gettoken()) != '\n') {
      if (type == PARENS || type == BRACKETS) {
        strcat(out, token);
      } else if (type == '*') {
        if ((type = nexttoken()) == PARENS || type == BRACKETS) {
          sprintf(temp, "(*%s)", out);
        } else {
          sprintf(temp, "*%s", out);
        }
        strcpy(out, temp);
      } else if (type == NAME) {
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

  if (prevtoken == YES) {
    prevtoken = NO;
    return tokentype;
  }

  while ((c = getch()) == ' ' || c == '\t') {
    ;
  }

  if (c == '(') {
    if ((c = getch()) == ')') {
      strcpy(token, "()");
      return tokentype = PARENS;
    } else {
      ungetch(c);
      return tokentype = '(';
    }
  } else if (c == '[') {
    for (*p++ = c; (*p++ = getch()) != ']';) {
      ;
    }

    *p = '\0';
    return tokentype = BRACKETS;
  } else if (isalpha(c)) {
    for (*p++ = c; isalnum(c = getch());) {
      *p++ = c;
    }
    *p = '\0';
    ungetch(c);
    return tokentype = NAME;
  } else {
    return tokentype = c;
  }
}

// get the next token and push it back
int nexttoken() {
  int type = gettoken();
  prevtoken = YES;

  return type;
}

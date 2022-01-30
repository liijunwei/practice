/*
page 109

修改dcl程序, 使它能够处理输入中的错误

TODO 没懂
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

enum { NO, YES };

void dcl();
void dirdcl();
void errmsg(char *msg);
int gettoken();
int tokentype;           /* 最后一个记号的类型    */
char token[MAXTOKEN];    /* 最后一个记号字符串    */
char name[MAXTOKEN];     /* 标识符名              */
char datatype[MAXTOKEN]; /* 数据类型为char/int 等 */
char out[1000];          /* 输出串                */
int prevtoken;

int main(int argc, char const *argv[])
{
  while(gettoken() != EOF) { /* 该行的第一个记号是数据类型 */
    strcpy(datatype, token);
    out[0] = '\0';
    dcl();

    if(tokentype != '\n') {
      errmsg("Syntax error\n");
    }

    printf("%s: %s %s\n", name, out, datatype);
  }

  return 0;
}

// parse a declarator
void dcl() {
  int ns;
  for(ns = 0; gettoken() == '*'; ) {
    ns++;
  }

  dirdcl();

  while(ns > 0) {
    strcat(out, " pointer to");
    ns--;
  }
}

// parse a direct declration
void dirdcl() {
  int type;

  if(tokentype == '(') {
    dcl();
    if(tokentype != ')') {
      errmsg("error: missing )\n");
    }
  } else if(tokentype == NAME) {
    strcpy(name, token);
  } else {
    errmsg("error: expected name of (dcl)\n");
  }

  while((type = gettoken()) == PARENS || type == BRACKETS) {
    if(type == PARENS) {
      strcat(out, " function returning");
    } else {
      strcat(out, " array");
      strcat(out, token);
      strcat(out, " of");
    }
  }
}

// print error message and indicate avail token
void errmsg(char *msg) {
  printf("%s", msg);
  prevtoken = YES;
}

/* 跳过空格和制表符 */
int gettoken() {
  int c;
  char *p = token;

  if(prevtoken == YES) {
    prevtoken = NO;
    return tokentype;
  }

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
    for(*p++ = c; (*p++ = getch()) != ']'; ) {
      ;
    }

    *p = '\0';
    return tokentype = BRACKETS;
  } else if(isalpha(c)) {
    for(*p++ = c; isalnum(c = getch()); ) {
      *p++ = c;
    }
    *p = '\0';
    ungetch(c);
    return tokentype = NAME;
  } else {
    return tokentype = c;
  }
}


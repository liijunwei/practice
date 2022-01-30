/*
page 107

根据声明符的语法对声明进行分析

TODO 补充注释

测试方法:
page 105
bash ch5-Pointers-and-Arrays/dcl-demo-test.sh
*/

#include <stdio.h>
#include <string.h>
#include <ctype.h>

#include "../common-utils/getch-ungetch.c"
#include "../common-utils/print-error.c"

#define MAXTOKEN 100

enum {
  NAME,
  PARENS,
  BRACKETS
};

void dcl();
void dirdcl();
int gettoken();
int tokentype;           /* 最后一个记号的类型    */
char token[MAXTOKEN];    /* 最后一个记号字符串    */
char name[MAXTOKEN];     /* 标识符名              */
char datatype[MAXTOKEN]; /* 数据类型为char/int 等 */
char out[1000];          /* 输出串                */

/* 将声明转为文字描述 */
/* 核心为两个函数: dcl和dirdcl */
/* 递归下降语法分析 */
int main(int argc, char const *argv[])
{
  while(gettoken() != EOF) { /* 该行的第一个记号是数据类型 */
    strcpy(datatype, token);
    out[0] = '\0';
    dcl();

    if(tokentype != '\n') {
      error("Syntax error\n");
    }

    printf("%s: %s %s\n", name, out, datatype);
  }

  return 0;
}

// dcl 和 dirdcl 相互递归调用
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

void dirdcl() {
  int type;

  if(tokentype == '(') {
    dcl();
    if(tokentype != ')') {
      error("error: missing )\n");
    }
  } else if(tokentype == NAME) {
    strcpy(name, token);
  } else {
    error("error: expected name of (dcl)\n");
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

/* 跳过空格和制表符 */
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


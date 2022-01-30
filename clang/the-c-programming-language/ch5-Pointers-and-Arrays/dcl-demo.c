/*
page 107

根据声明符的语法对声明进行分析

TODO 补充注释

测试方法:
page 105
echo "char **argv"       | ./a.out
echo "int (*daytab)[13]" | ./a.out
echo "int *daytab[13]"   | ./a.out
echo "void *comp()"      | ./a.out
echo "void (*comp)()"    | ./a.out
echo "(*(*x())[])()"     | ./a.out # TODO
echo "(*(*x[3])())[5]"   | ./a.out # TODO
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
/* 核心为两个函数: dcl和dirdcl */
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

// dcl 和 dirdcl 相互递归调用
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
  int type;

  if(tokentype == '(') {
    dcl();
    if(tokentype != ')') {
      printf("error: missing )\n");
    }
  } else if(tokentype == NAME) {
    strcpy(name, token);
  } else {
    printf("error: expected name of (dcl)\n");
  }

  while((type = gettoken()) == PARENS || type == BRACKETS) {
    if(type == PARENS) {
      strcat(out, " function returning");
    } else {
      strcat(out, " array");
      strcat(out, "token");
      strcat(out, " of");
    }
  }
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


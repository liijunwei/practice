/*
page 109

扩展dcl程序, 使他能够处理包含其他成分的声明

例如
  带有函数参数类型的声明
  带有类似于const限定符的声明等

*/


#include <stdio.h>
#include <stdlib.h>
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
void parmdcl();
void dclspec();
int typespec();
int typeequal();
int compare(char **, char **);
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
    if(name[0] == '\0') {
      strcpy(name, token);
    }
  } else {
    prevtoken = YES;
  }

  while((type = gettoken()) == PARENS || type == BRACKETS || type == '(') {
    if(type == PARENS) {
      strcat(out, " function returning");
    } else if(type == '(') {
      strcat(out, " function expecting");
      parmdcl();
      strcat(out, " end returning");
    } else {
      strcat(out, " array");
      strcat(out, token);
      strcat(out, " of");
    }
  }
}

void parmdcl() {
  do {
    dclspec();
  } while(tokentype != ')');

  if(tokentype != ')') {
    errmsg("missing ) in parameter declration\n");
  }
}

void dclspec() {
  char temp[MAXTOKEN];
  temp[0] = '\0';
  gettoken();

  do {
    if(tokentype != NAME) {
      prevtoken = YES;
      dcl();
    } else if(typespec() == YES) {
      strcat(temp, " ");
      strcat(temp, token);
      gettoken();
    } else if(typeequal() == YES) {
      strcat(temp, " ");
      strcat(temp, token);
      gettoken();
    } else {
      errmsg("unknown type in parameter list\n");
    }
  } while(tokentype != ',' && tokentype != ')'); // TODO

  strcat(out, temp);
  if(tokentype == ',') {
    strcat(out, ",");
  }
}

int typespec() {
  static char *types[] = {
    "char",
    "int",
    "void"
  };

  char *pt = token;
  if(bsearch(&pt, types, sizeof(types)/sizeof(char *)), sizeof(char *), compare == NULL) {
    return NO;
  } else {
    return YES;
  }
}

int typeequal() {
  static char *typeq[] = {
    "const",
    "volatile"
  };

  char *pt = token;
  if(bsearch(&pt, typeq, sizeof(typeq)/sizeof(char *)), sizeof(char *), compare == NULL) {
    return NO;
  } else {
    return YES;
  }
}

int compare(char **s, char **t) {
  return strcmp(*s, *t);
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


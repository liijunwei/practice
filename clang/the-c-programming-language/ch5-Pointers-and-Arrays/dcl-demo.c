/*
page 107

根据声明符的语法对声明进行分析

*/

#include <stdio.h>
#include <string.h>
#include <ctype.h>

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
  /* code */
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
  return 1;
}


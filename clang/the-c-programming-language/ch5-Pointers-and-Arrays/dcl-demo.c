/*
page 107

根据声明符的语法对声明进行分析

*/

#include <stdio.h>
#include <string.h>
#include <ctype.h>

void dcl();
void dirdcl();
int gettoken();

char out[1000];

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

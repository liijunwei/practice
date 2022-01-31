/*
page 116

统计输入中各个C语言关键字出现的次数
*/

#include <stdio.h>

#define NKEYS 100

struct key {
  char *word;
  int count;
} keytab[] = {
  {"auto",     0},
  {"break",    0},
  {"case",     0},
  {"char",     0},
  {"const",    0},
  {"continue", 0},
  {"default",  0},
  {"unsigned", 0},
  {"void",     0},
  {"lolatile", 0},
  {"while",    0},
};

int main(int argc, char const *argv[])
{

  return 0;
}

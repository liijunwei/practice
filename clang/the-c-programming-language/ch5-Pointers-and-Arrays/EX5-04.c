#include <stdio.h>
#include <assert.h>
#include <string.h>

/*
page 92

编写函数strend(s, t) 如果字符串t出现在字符串s的尾部, 该函数返回1, 否则返回0

*/

int strend(char *s, char *t){
  int len_s = strlen(s);
  int len_t = strlen(t);

  if(len_t == 0 || (len_t > len_s)){
    return 0;
  }

  return 1;
}

int main(int argc, char const *argv[])
{

  char str[100] = "Hello";

  assert(1 == strend(str, "llo"));
  assert(0 == strend(str, "Helloooo"));
  assert(0 == strend(str, ""));

  return 0;
}

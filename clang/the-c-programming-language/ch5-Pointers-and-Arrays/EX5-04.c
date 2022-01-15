#include <stdio.h>
#include <assert.h>
#include <string.h>

/*
page 92

编写函数strend(s, t) 如果字符串t出现在字符串s的尾部, 该函数返回1, 否则返回0

*/

int strend(char *s, char *t){
  char *begin_s = s;
  char *begin_t = t;

  for(; *s != '\0'; s++){
    ;
  }

  for(; *t != '\0'; t++){
    ;
  }

  for(; *s == *t; s--, t--){
    if(t == begin_t || s == begin_s){
      break;
    }
  }

  if(*s == *t && t == begin_t && *s != '\0'){
    return 1;
  } else {
    return 0;
  }
}

int main(int argc, char const *argv[])
{

  char str[100] = "Hello";

  assert(1 == strend(str, "Hello"));
  assert(1 == strend(str, "llo"));
  assert(1 == strend(str, "o"));
  assert(0 == strend(str, "hei"));
  assert(0 == strend(str, "Helloooo"));
  assert(0 == strend(str, ""));

  return 0;
}

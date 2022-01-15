#include <stdio.h>
#include <assert.h>
#include <string.h>

/*
page 92

编写函数strend(s, t) 如果字符串t出现在字符串s的尾部, 该函数返回1, 否则返回0

*/

// TODO 没看懂...
int strend_v1(char *s, char *t){
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

// https://github.com/fool2fish/the-c-programming-language-exercise-answers/blob/master/ch05/5-4-strend/strend.c
int strend_v2(char *s, char *t) {
  int len_s = strlen(s);
  int len_t = strlen(t);

  if(len_t == 0 || len_s < len_t){
    return 0;
  }

  while(len_t > 1) {
    if(*(t + len_t) != *(s + len_s)) {
      return 0;
    }

    len_t--;
    len_s--;
  }

  return 1;
}

int main(int argc, char const *argv[])
{

  char str[100] = "Hello";

  assert(1 == strend_v1(str, "Hello"));
  assert(1 == strend_v1(str, "llo"));
  assert(1 == strend_v1(str, "o"));
  assert(0 == strend_v1(str, "hei"));
  assert(0 == strend_v1(str, "Helloooo"));
  assert(0 == strend_v1(str, ""));

  assert(1 == strend_v2(str, "Hello"));
  assert(1 == strend_v2(str, "llo"));
  assert(1 == strend_v2(str, "o"));
  assert(0 == strend_v2(str, "hei"));
  assert(0 == strend_v2(str, "Helloooo"));
  assert(0 == strend_v2(str, ""));

  return 0;
}

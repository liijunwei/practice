#include <stdio.h>
#include <assert.h>

/*
page 85

求字符串的长度
*/

int custom_strlen(const char *s){
  int n;

  for(n = 0; *s != '\0'; s++){
    n++;
  }

  return n;
}

int main(int argc, char const *argv[])
{

  assert(0 == custom_strlen(""));
  assert(1 == custom_strlen("a"));
  assert(2 == custom_strlen("ab"));
  assert(3 == custom_strlen("abc"));
  assert(4 == custom_strlen("abcd"));

  return 0;
}


#include <stdio.h>
#include <string.h>
#include <assert.h>

/*
page 75

编写一个递归版本的reverse(s)函数, 以将字符串s倒置

*/

void reverse(char s[]){
  int c;
  int i;
  int j;

  for(i = 0, j = strlen(s) - 1; i < j; i++, j--){
    c = s[i];
    s[i] = s[j];
    s[j] = c;
  }
}

#define ARR_SIZE 10

int main(int argc, char const *argv[])
{
  char str[ARR_SIZE] = "abcdefg";

  reverse(str);
  assert(strcmp("gfedcba", str) == 0);

  return 0;
}

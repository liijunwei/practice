#include <stdio.h>
#include <string.h>
#include <assert.h>

/*
page 53

*/

// 删除字符串尾部的空格符/制表符和换行符
int trim(char s[]){
  int n;

  for(n = strlen(s) - 1; n >= 0; n--){
    if(s[n] != ' ' && s[n] != '\t' && s[n] != '\n'){
      break;
    }

  }

  s[n+1] = '\0';

  return n;
}

int main(int argc, char const *argv[])
{
  char str[100] = "hello \n 20211230 \t :D";

  assert(100 == trim(str));
  printf("PASS.\n");

  return 0;
}
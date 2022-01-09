#include <stdio.h>
#include <ctype.h>
#include <assert.h>

/*
page 50

*/

// 将s转化为整型数
int atoi(char s[]){
  int i;
  int n;
  int sign;

  for(i = 0; isspace(s[i]); i++){
    ; // 跳过空白符
  }

  sign = (s[i] == '-') ? -1 : 1;

  if(s[i] == '+' || s[i] == '-'){
    i++;
  }

  for(n = 0; isdigit(s[i]); i++){
    n = 10 * n + (s[i] - '0');
  }

  return sign * n;
}

int main(int argc, char const *argv[])
{
  assert(12 == atoi("+12"));
  assert(12 == atoi("12"));
  assert(120 == atoi("120"));
  assert(-120 == atoi("-120"));
  printf("PASS.\n");

  return 0;
}

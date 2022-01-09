#include <stdio.h>
#include <ctype.h>
#include <assert.h>

/*
page 62

利用atof把字符串s转换为整数
*/

// 把字符串s转换为相应的双精度浮点数
double atof(const char s[]){
  double val;
  double power;

  int i;
  int sign;

  for(i = 0; isspace(s[i]); i++){
    ; // 跳过空白符
  }

  sign = (s[i] == '-') ? -1 : 1;

  if(s[i] == '+' || s[i] == '-'){
    i++;
  }

  for(val = 0.0; isdigit(s[i]); i++){
    val = 10.0 * val + (s[i] - '0');
  }

  if(s[i] == '.'){
    i++;
  }

  for(power = 1.0; isdigit(s[i]); i++){
    val = 10.0 * val + (s[i] - '0');
    power *= 10.0;
  }

  return sign * val / power;
}

int atoi(const char s[]){
  return (int)atof(s);
}

int main(int argc, char const *argv[])
{
  char str1[100] = "314159";
  assert(314159 == atoi(str1));

  char str2[100] = "123730951";
  assert(123730951 == atoi(str2));

  char str3[100] = "889";
  assert(889 == atoi(str3));

  printf("PASS.");
  return 0;
}

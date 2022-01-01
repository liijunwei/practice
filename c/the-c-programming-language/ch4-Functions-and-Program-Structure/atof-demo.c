#include <stdio.h>
#include <ctype.h>
#include <assert.h>

/*
page 60

标准库<stdlib.h>里实现了atof函数, 比这个版本好

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

int main(int argc, char const *argv[])
{
  char str1[100] = "3.14159265354"; // PI
  assert(3.14159265354 == atof(str1));

  char str2[100] = "1.4142135623730951"; // Sqrt(2)
  assert(1.4142135623730951 == atof(str2));

  char str3[100] = "0.6931471805599453"; // log(2)
  assert(0.6931471805599453 == atof(str3));

  printf("PASS.");
  return 0;
}
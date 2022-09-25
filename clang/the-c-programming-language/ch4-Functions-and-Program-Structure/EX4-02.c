#include <assert.h>
#include <ctype.h>
#include <stdio.h>

/*
page 62

对 atof 函数进行扩充, 使他可以处理形如 123.45e-6 的科学表示法
其中, 浮点数后可能会紧跟一个e或E以及一个指数(可能有正负号)
*/

// 把字符串s转换为相应的双精度浮点数
double atof(const char s[]) {
  double val;
  double power;

  int i;
  int sign;
  int exp;

  for (i = 0; isspace(s[i]); i++) {
    ; // 跳过空白符
  }

  sign = (s[i] == '-') ? -1 : 1;

  if (s[i] == '+' || s[i] == '-') {
    i++;
  }

  for (val = 0.0; isdigit(s[i]); i++) {
    val = 10.0 * val + (s[i] - '0');
  }

  if (s[i] == '.') {
    i++;
  }

  for (power = 1.0; isdigit(s[i]); i++) {
    val = 10.0 * val + (s[i] - '0');
    power *= 10.0;
  }

  val = sign * val / power;

  if (tolower(s[i]) == 'e') {
    sign = (s[++i] == '-') ? -1 : 1;

    if (s[i] == '+' || s[i] == '-') {
      i++;
    }

    for (exp = 0; isdigit(s[i]); i++) {
      exp = 10 * exp + (s[i] - '0');
    }

    if (sign == 1) {
      while (exp > 0) { // positive exponent
        val *= 10;
        exp--;
      }
    } else {
      while (exp > 0) { // negative exponent
        val /= 10;
        exp--;
      }
    }
  }

  return val;
}

int main(int argc, char const *argv[]) {
  char str1[100] = "1.23e2";
  assert(123 == atof(str1));

  char str2[100] = "1.23e-2";
  assert(0.0123 == atof(str2));

  char str3[100] = "10e3";
  assert(10000 == atof(str3));

  char str4[100] = "10E-3";
  assert(0.01 == atof(str4));

  return 0;
}

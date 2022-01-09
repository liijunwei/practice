#include <stdio.h>
#include <ctype.h>

/*
page 61

计算器程序
每行中读取一个数(数前面可能有正负号), 并对它们求和, 在每次输入完成后把这些书的累积总和打印出来
*/

#define MAXLINE 1000 // 输入行的最大长度

// 将行保存到s中, 并返回该行的行数
int custom_getline(char s[], int max){
  int c;
  int i = 0;

  while(--max > 0 && (c = getchar()) != EOF && c != '\n'){
    s[i] = c;
    i++;
  }

  if(c == '\n'){
    s[i] = c;
    ++i;
  }

  s[i] = '\0';

  return i;
}

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
  double sum = 0;
  char line[MAXLINE];

  while(custom_getline(line, MAXLINE) > 0){
    printf("sum is: %g\n", sum += atof(line));
  }

  return 0;
}

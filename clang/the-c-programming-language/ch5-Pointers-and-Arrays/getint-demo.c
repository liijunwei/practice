#include <stdio.h>
#include <ctype.h>

#include "../common-utils/getch-ungetch.c"
#include "../common-utils/print-array.c"

/*
page 82

函数getint接收自由格式的输入, 并执行转换, 将输入的字符流分解为整数, 且每次调用得到一个整数
getint需要返回转换后得到的整数, 并且在到达输入结尾时要返回文件结束标记
这些值必须以不同的方式返回
EOF可以用任何值表示, 当然也可用一个输入的整数表示

*/

#define ARR_SIZE 10

int getint(int *pn){
  int c;
  int sign;

  while(isspace(c = getch())){
    ;
  }

  if(!isdigit(c) && c != EOF && c != '+' && c != '-'){
    ungetch(c); // 输入不是数字
    return 0;
  }

  sign = (c == '-') ? -1 : 1;
  if(c == '+' || c == '-'){
    c = getch();
  }

  for(*pn = 0; isdigit(c); c = getch()){
    *pn = 10 * (*pn) + (c - '0');
  }

  *pn *= sign;

  if(c != EOF){
    ungetch(c);
  }

  return c;
}

int main(int argc, char const *argv[])
{
  int n;
  int array[ARR_SIZE];

  for (n = 0; n < ARR_SIZE; n++){
    array[n] = 0;
  }
  print_int_array(array, ARR_SIZE);

  for(n = 0; n < ARR_SIZE && getint(&array[n]) != EOF; n++){
    ;
  }
  print_int_array(array, ARR_SIZE);

  return 0;
}

#include <stdio.h>
#include <ctype.h>

#include "../common-utils/getch-ungetch.c"
#include "../common-utils/print-array.c"

/*
page 83

模仿getint, 写一个读取浮点数的函数getfloat
getfloat的返回值应该是什么类型?

*/

int getfloat(float *pn){
  int c;
  int sign;
  float power;

  while(isspace(c = getch())){
    ;
  }

  // not a number
  if(!isdigit(c) && c != EOF && c != '+' && c != '-' && c != '.'){
    ungetch(c);
    return 0;
  }

  sign = (c == '-') ? -1 : 1;

  if(c == '+' || c == '-'){
    c = getch();
  }

  for(*pn = 0.0; isdigit(c); c = getch()){
    *pn = 10.0 * (*pn) + (c - '0'); // integer part
  }

  if(c == '.'){
    c = getch();
  }

  for (power = 1.0; isdigit(c); c = getch()){
    *pn = 10.0 * (*pn) + (c - '0'); // fractional part
    power *= 10.0;
  }

  *pn = (*pn) * sign / power;

  if(c != EOF){
    ungetch(c);
  }

  return c;
}

#define ARR_SIZE 10

int main(int argc, char const *argv[])
{
  int n;
  float array[ARR_SIZE];

  for (n = 0; n < ARR_SIZE; n++){
    array[n] = 0;
  }

  for(n = 0; n < ARR_SIZE && getfloat(&array[n]) != EOF; n++){
    ;
  }
  print_float_array(array, ARR_SIZE);

  return 0;
}

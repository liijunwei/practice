#include <stdio.h>
#include <ctype.h>

#include "../common-utils/getch-ungetch.c"
#include "../common-utils/print-array.c"

/*
page 83

ch5-Pointers-and-Arrays/getint-demo.c
上面的例子中,如果符号 '+' '-' 后面紧跟的不是数字, getint函数会把符号视为数字0的有效表达式
修改该函数,将这种+ - 符号重新协会到输入流中

OK 还是不太明白 getch ungetch 输入流 "压回" 是什么意思...
TODO 没明白为什么后面紧跟的不是数字的时候, 后面的数字被视为0

*/

#define ARR_SIZE 10

int getint(int *pn){
  int c;
  int d;
  int sign;

  while(isspace(c = getch())){
    ;
  }

  if(!isdigit(c) && c != EOF && c != '+' && c != '-'){
    ungetch(c); // 输入不是数字
    return 0;
  }

  sign = (c == '-') ? -1 : 1;
  if(c == '+' || c == '-'){ // only run once if the input is appropriate
    d = c; // remember sign char
    if(!isdigit(c = getch())){
      if(c != EOF){
        ungetch(c); // push back non-digit
      }

      ungetch(d);  // push back sign char
      return d;    // return sign char
    }
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

  for(n = 0; n < ARR_SIZE && getint(&array[n]) != EOF; n++){
    ;
  }
  print_int_array(array, ARR_SIZE);

  return 0;
}

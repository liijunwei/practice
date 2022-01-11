#include <stdio.h>
#include <ctype.h>

/*
page 83

ch5-Pointers-and-Arrays/getint-demo.c
上面的例子中,如果符号 '+' '-' 后面紧跟的不是数字, getint函数会把符号视为数字0的有效表达式
修改该函数,将这种+ - 符号重新协会到输入流中


*/

#define ARR_SIZE 10
#define BUFSIZE  100

char buf[BUFSIZE];
int bufp = 0; // buf中下一空闲的位置

// 取一个字符(可能是压回的字符)
int getch(void){
  return (bufp > 0) ? buf[--bufp] : getchar();
}

// 把字符压回输入中
void ungetch(int c){
  if(bufp >= BUFSIZE) {
    printf("Ungetch, too many characters\n");
  } else {
    buf[bufp++] = c;
  }
}


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
  if(c == '+' || c == '-'){
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

void print_arr(int a[]){
  for(int i = 0; i < ARR_SIZE; i++){
    printf("%d ", a[i]);
  }
  printf("\n");
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
  print_arr(array);

  return 0;
}

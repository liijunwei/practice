#include <stdio.h>
#include <ctype.h>

/*
page 82

函数getint接收自由格式的输入, 并执行转换, 将输入的字符流分解为整数, 且每次调用得到一个整数
getint需要返回转换后得到的整数, 并且在到达输入结尾时要返回文件结束标记
这些值必须以不同的方式返回
EOF可以用任何值表示, 当然也可用一个输入的整数表示

TODO not clear
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

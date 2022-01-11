#include <stdio.h>

/*
理解 getch 和 ungetch函数

*/

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

void print_buf(char c[]){
  for(char i = 0; i < BUFSIZE; i++){
    printf("%c ", c[i]);
  }
  printf("\n");
}

int main(int argc, char const *argv[])
{
  ungetch('1');
  ungetch('2');
  ungetch('3');

  print_buf(buf);
  putchar(getch());
  putchar(getch());
  putchar(getch());

  putchar(getch());
  putchar(getch());
  putchar(getch());

  return 0;
}

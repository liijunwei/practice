#include <stdio.h>
#include <limits.h>

/*
page 33

for(int i = 0; i < lim -1 && (c = getchar()) != '\n' && c != EOF; ++i){ s[i] = c;}
在不使用&&或||的条件下编写一个与上面的for循环语句等价的循环语句
*/

#define FLAG_TRUE   1
#define FLAG_FALSE -1

int main(int argc, char const *argv[])
{
  int flag = FLAG_TRUE;
  int i = 0;
  int limit = 10;
  char s[limit];
  int c;

  while(flag == FLAG_TRUE){
    if(i > limit - 1){
      flag = FLAG_FALSE;
    } else if((c = getchar()) == '\n'){
      flag = -1;
    } else if(c == EOF){
      flag = -1;
    } else{
      s[i] = c;
      ++i;
    }
  }

  printf("str is: %s\n", s);

  return 0;
}

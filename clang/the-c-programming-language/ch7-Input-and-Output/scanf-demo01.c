/*
page 137

输入函数scanf对应输入函数printf, 他在与后者相反的方向上提供同样的转换功能
`int scanf(char *fmt, ...)`
*/

#include <stdio.h>

// gcc -g ch7-Input-and-Output/scanf-demo01.c && echo 20 "Monkey-D-Luffy" | ./a.out

int main(int argc, char const *argv[])
{
  int number;
  printf("Please enter a number: \n");
  scanf("%d", &number);
  printf("You entered: %d\n", number);

  char name[100];
  printf("Please enter a name: \n");
  scanf("%s", name);
  printf("You entered: %s\n", name);

  return 0;
}

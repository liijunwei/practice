/*
page 139

读入包含下列包含日期格式的输入行
25 Dec 1988
*/

#include <stdio.h>

// gcc -g ch7-Input-and-Output/scanf-demo03.c && echo "25 Dec 1988" | ./a.out
int main(int argc, char const *argv[])
{
  int day;
  int year;

  char monthname[20];

  scanf("%d %s %d", &day, monthname, &year);

  printf("year:      %d\n", year);
  printf("monthname: %s\n", monthname);
  printf("day:       %d\n", day);

  return 0;
}

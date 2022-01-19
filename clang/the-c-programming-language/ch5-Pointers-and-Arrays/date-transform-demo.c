#include <stdio.h>

/*
page 95

把某月某日这种日期形式转化为某年中的第几天的表示形式
反之亦然
*/

static char daytab[2][13] = {
  {0, 31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31},
  {0, 31, 29, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
};

// 将某月某日的日期表示形式转换为某年中第几天的表现形式
int day_of_year(int year, int month, int day){
  int i;
  int leap;

  leap = year % 4 && year % 100 != 0 || year % 400;

  for(i = 1; i < month; i++){
    day += daytab[leap][i];
  }

  return day;
}

int main(int argc, char const *argv[])
{

  return 0;
}

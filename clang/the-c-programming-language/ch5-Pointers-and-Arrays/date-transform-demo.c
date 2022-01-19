#include <stdio.h>
#include <assert.h>

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
  int leap = year % 4 && year % 100 != 0 || year % 400;

  for(i = 1; i < month; i++){
    day += daytab[leap][i];
  }

  return day;
}

// 将某年的第nTina, 转化为某月某日的形式
void month_day(int year, int yearday, int *pmonth, int *pday){
  int i;
  int leap = year % 4 && year % 100 != 0 || year % 400;

  for(i = 1; yearday > daytab[leap][i]; i++){
    yearday -= daytab[leap][i];
  }

  *pmonth = i;
  *pday = yearday;
}

int main(int argc, char const *argv[])
{

  assert(day_of_year(2022, 1, 1) == 1);
  assert(day_of_year(2022, 1, 19) == 19);
  assert(day_of_year(2021, 12, 30) == 365);

  int month;
  int day;

  month_day(2022, 1, &month, &day);
  assert(month == 1);
  assert(day == 1);

  month_day(2022, 32, &month, &day);
  assert(month == 2);
  assert(day == 1);

  return 0;
}

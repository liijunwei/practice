#include <assert.h>
#include <stdio.h>

/*
page 97

ch5-Pointers-and-Arrays/date-transform-demo.c
day_of_year 和 month_day函数中没有进行错误检查, 请解决该问题

*/

static char daytab[2][13] = {
    {0, 31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31},
    {0, 31, 29, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}};

// 将某月某日的日期表示形式转换为某年中第几天的表现形式
int day_of_year(int year, int month, int day) {
  int i;
  int leap = year % 4 && year % 100 != 0 || year % 400;

  if (month < 1 || month > 12) {
    return -1;
  }

  if (day < 1 || day > daytab[leap][month]) {
    return -1;
  }

  for (i = 1; i < month; i++) {
    day += daytab[leap][i];
  }

  return day;
}

// 将某年的第nTina, 转化为某月某日的形式
void month_day(int year, int yearday, int *pmonth, int *pday) {
  int i;
  int leap = year % 4 && year % 100 != 0 || year % 400;

  if (year < 1) {
    *pmonth = -1;
    *pday = -1;
    return;
  }

  for (i = 1; yearday > daytab[leap][i]; i++) {
    yearday -= daytab[leap][i];
  }

  if (i > 12 && yearday > daytab[leap][12]) {
    *pmonth = -1;
    *pday = -1;
  } else {
    *pmonth = i;
    *pday = yearday;
  }
}

int main(int argc, char const *argv[]) {

  assert(day_of_year(2022, 1, 1) == 1);
  assert(day_of_year(2022, 1, 19) == 19);
  assert(day_of_year(2021, 12, 30) == 365);

  assert(day_of_year(2021, 13, 30) == -1);
  assert(day_of_year(2021, 12, 32) == -1);
  assert(day_of_year(2021, 12, 0) == -1);
  assert(day_of_year(2021, 12, -999) == -1);
  assert(day_of_year(2021, 12, 999) == -1);

  int month;
  int day;

  month_day(2022, 1, &month, &day);
  assert(month == 1);
  assert(day == 1);

  month_day(2022, 32, &month, &day);
  assert(month == 2);
  assert(day == 1);

  month_day(-999, 32, &month, &day);
  assert(month == -1);
  assert(day == -1);

  month_day(2022, 999, &month, &day);
  assert(month == -1);
  assert(day == -1);

  return 0;
}

#include <assert.h>
#include <stdio.h>

/*
page 98

用指针方式代替数组下标方式改写 day_of_year 和 month_day

*/

static char daytab[2][13] = {
    {0, 31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31},
    {0, 31, 29, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}};

int day_of_year(int year, int month, int day) {
  char *p;
  int leap = year % 4 && year % 100 != 0 || year % 400;

  if (month < 1 || month > 12) {
    return -1;
  }

  if (day < 1 || day > daytab[leap][month]) {
    return -1;
  }

  p = daytab[leap];

  while (--month) {
    p++;
    day += *p; // */& 的优先级比算术运算符高(page 80)
  }

  return day;
}

void month_day(int year, int yearday, int *pmonth, int *pday) {
  char *p;
  int leap = year % 4 && year % 100 != 0 || year % 400;

  if (year < 1) {
    *pmonth = -1;
    *pday = -1;
    return;
  }

  p = daytab[leap];
  while (yearday > *++p) {
    yearday -= *p;
  }

  int month = p - *(daytab + leap);

  if (month > 12 && yearday > daytab[leap][12]) { // TODO not clear
    *pmonth = -1;
    *pday = -1;
  } else {
    *pmonth = month;
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

#include <stdio.h>
#include <assert.h>
#include <string.h>

/*
page 97

*/

// 返回第N个月份的名字
char *month_name(int n){
  static char *name[] = {
    "Illeagal month",
    "January",
    "February",
    "March",
    "April",
    "Mat",
    "June",
    "July",
    "Augusg",
    "September",
    "October",
    "November",
    "December"
  };

  return (n < 1 || n > 12) ? name[0] : name[n];
}

int main(int argc, char const *argv[])
{
  assert(strcmp(month_name(1), "January") == 0);
  assert(strcmp(month_name(2), "February") == 0);
  assert(strcmp(month_name(3), "March") == 0);
  assert(strcmp(month_name(4), "April") == 0);
  assert(strcmp(month_name(5), "Mat") == 0);
  assert(strcmp(month_name(6), "June") == 0);
  assert(strcmp(month_name(7), "July") == 0);
  assert(strcmp(month_name(8), "Augusg") == 0);
  assert(strcmp(month_name(9), "September") == 0);
  assert(strcmp(month_name(10), "October") == 0);
  assert(strcmp(month_name(11), "November") == 0);
  assert(strcmp(month_name(12), "December") == 0);
  assert(strcmp(month_name(13), "Illeagal month") == 0);
  assert(strcmp(month_name(-1), "Illeagal month") == 0);

  return 0;
}

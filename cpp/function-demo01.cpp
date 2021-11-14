#include <iostream>

using namespace std;

/*
全局变量 破坏了函数的 相对独立性
全局变量 增加了函数之间的 耦合性
全局变量 使得函数之间的 交互不够清晰

建议: 在非必要的情况下, 不使用全局变量
*/


/*
  问题描述:

  给定从公元2000-01-01开始过去的天数, 请编写程序给出这一天是哪年哪月哪日星期几
  注意: 闰年被定义为能被4整除的年份, 但是能被100整除而不能被400整除的年是例外, 他们不是闰年;
    例如: 1700 1800 1900 2100 不是闰年
    例如: 1600 2000 2400 是闰年
  注意: 每个月的天数不同
  注意: 假设结果的年份不会超过9999

  输入要求:

  输入多组数据, 每组一个正整数, 表示从2000-1-1开始已过去的天数
  对输入的每个天数, 输出一行, 格式为 "YYYY-MM-DD DayOfWeek", 其中 DayOfWeek 必须是 Sunday, Monday...
  输出最后一行是-1, 不必处理
*/

/*
  思路:

  1. 算出星期几
  2. 减掉每年的天数 -> year
  3. 减掉每月的天数 -> month
  4.                -> day
*/

int days;
int get_dayofweek();
int get_year();
int get_month(int leap_year);
int get_leap_year(int y);

int main(){
  int year;
  int month;
  int dayofweek;
  int leap_year;

  char week[7][10] = {
    "Saturday",
    "Sunday",
    "Monday",
    "Tuesday",
    "Wednesday",
    "Thursday",
    "Friday",
  };

  while((cin >> days) && days != -1){
    dayofweek = get_dayofweek();
    year      = get_year();
    leap_year = get_leap_year(year);
    month     = get_month(leap_year);

    cout << year  << "-"
         << month << "-"
         << days  << "-"
         << week[dayofweek] << endl;

   cout << endl;
  }

  return 0;
}

int get_dayofweek(){
  return days % 7;
}

int get_leap_year(int y){
  return (y % 4 == 0 && y % 100 != 0 || y % 400 == 0);
}

int get_year(){
  int i = 2000;
  int leap_year;
  while(true){
    leap_year = get_leap_year(i);
    if(leap_year == 1 && days >= 366){
      days = days - 366;
      i++;
      continue;
    }
    else if(leap_year == 0 && days > 365){
      days = days - 365;
      i++;
      continue;
    }
    else{
      break;
    }
  }

  return i;
}

int get_month(int leap_year){
  int pmonth[12] = {31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31};
  int rmonth[12] = {31, 29, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31};

  int j = 0;

  while(true){
    if(leap_year == 1 && days >= rmonth[j]){
      days = days - rmonth[j];
      j++;
    }
    else if(leap_year == 0 && days >= pmonth[j]){
      days = days - pmonth[j];
      j++;
    }
    else{
      break;
    }
  }

  return ++j;
}


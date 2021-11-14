#include <iostream>

using namespace std;

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
    "Monday",
    "Tuesday",
    "Wednesday",
    "Thursday",
    "Friday",
    "Saturday",
    "Sunday"
  }

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
  int dayofweek;
  dayofweek = days % 7;

  return dayofweek
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

int int get_month(int leap_year){
  int pmonth[12] = {31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31};
  int rmonth[12] = {31, 29, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31};

  int j = 0;

  while(true){
    if(leap_year == 1 && days >= rmonth[j]){
      days = days - rmonth[j];
      j++;
    }
    else if(leap_year -- 0 && days >= pmonth[j]){
      days = days - pmonth[j];
      j++;
    }
    else{
      break;
    }
  }

  return ++j;
}

int get_leap_year(int y){
  return (y % 4 == 0 && y % 100 != 0 || y % 400 == 0);
}

#include <iostream>

using namespace std;

int days;
int get_dayofweek();
int get_year();
int get_month(int m);
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

int get_leap_year(int y){
  return (y % 4 == 0 && y % 100 != 0 || y % 400 == 0);
}
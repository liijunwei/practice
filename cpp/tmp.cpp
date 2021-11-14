#include <iostream>

using namespace std;

int get_leap_year(int y){
  return (y % 4 == 0 && (y % 100 != 0 || y % 400 == 0));
}

int main(){
  cout << "get_leap_year(1700) " << get_leap_year(1700) << endl;
  cout << "get_leap_year(1800) " << get_leap_year(1800) << endl;
  cout << "get_leap_year(1900) " << get_leap_year(1900) << endl;
  cout << "get_leap_year(2100) " << get_leap_year(2100) << endl;

  cout << "get_leap_year(1600) " << get_leap_year(1600) << endl;
  cout << "get_leap_year(2000) " << get_leap_year(2000) << endl;
  cout << "get_leap_year(2400) " << get_leap_year(2400) << endl;
  cout << endl;
  return 0;
}

#include <iostream>

using namespace std;

int get_leap_year(int y){
  return (y % 4 == 0 && y % 100 != 0 || y % 400 == 0);
}

int main(){

  cout << "get_leap_year(2000) " << get_leap_year(2000) << endl;
  cout << "get_leap_year(2001) " << get_leap_year(2001) << endl;
  cout << "get_leap_year(2002) " << get_leap_year(2002) << endl;
  cout << endl;
  return 0;
}

#include <iostream>

using namespace std;

int main(){
  char a[3][40] = {
    "0 -> element-0",
    "1 -> element-1",
    "2 -> element-2",
  };

  cout << __LINE__ << " " << a[0] << endl;
  cout << __LINE__ << " " << a[1] << endl;
  cout << __LINE__ << " " << a[2] << endl;
  cout << endl;

  cout << __LINE__ << " " << *a[0] << endl;
  cout << __LINE__ << " " << *a[1] << endl;
  cout << __LINE__ << " " << *a[2] << endl;
  cout << endl;

  cout << endl;
  return 0;
}

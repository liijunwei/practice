#include <iostream>

using namespace std;

int main(){
  int iCount = 18;
  cout << "iCount  " << iCount << endl;

  int *iPtr = &iCount;
  *iPtr = 58;

  cout << "iCount  " << iCount << endl;
  cout << "iPtr    " << iPtr << endl;
  cout << "&iCount " << &iCount << endl;
  cout << "*iPtr   " << *iPtr << endl;
  cout << "&iPtr   " << &iPtr << endl;

  cout << endl;
  return 0;
}

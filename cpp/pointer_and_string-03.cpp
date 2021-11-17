#include <iostream>

using namespace std;

int main(){
  char buffer[10] = "ABC";
  char *pc;

  pc = "hello";
  cout << __LINE__ << " " << pc << endl;
  pc++;

  cout << __LINE__ << " " << pc << endl;
  cout << __LINE__ << " " << *pc << endl;

  pc = buffer;
  cout << __LINE__ << " " << pc << endl;

  cout << endl;
  return 0;
}

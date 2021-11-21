#include <iostream>

using namespace std;

int main(){
  char buffer[10] = "ABC";
  const char *pc;

  // Important note:
  // we can assign "hello" to pointer pc
  // but we cannot change "hello" by pc, cause "hello" is stored as constant
  // we're not allowed to change a constant.
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

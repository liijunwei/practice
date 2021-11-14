#include <iostream>

using namespace std;

// 切饼问题

int cut(int n);

int main(){
  cout << "cut(1): "  << cut(1)  << endl;
  cout << "cut(2): "  << cut(2)  << endl;
  cout << "cut(3): "  << cut(3)  << endl;
  cout << "cut(4): "  << cut(4)  << endl;
  cout << "cut(5): "  << cut(5)  << endl;
  cout << "cut(50): " << cut(50) << endl;
  cout << "cut(60): " << cut(60) << endl;
  cout << "cut(70): " << cut(70) << endl;

  cout << endl;
  return 0;
}

int cut(int n){
  if(n == 0){
    return 1;
  }
  else{
    return (n + cut(n - 1));
  }
}
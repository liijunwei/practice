#include <iostream>

using namespace std;

// variable.1 `excel_number` in global scope
// variable.2 `socre`        in excel_count function
// variable.3 `score`        in main function
// variable.4 `i`            in for loop

int excel_number = 0;

void excel_count(float socre);

int main(){
  float score = 0;

  for(int i = 0; i < 3; i++){
    cin >> score;
    excel_count(score);
  }

  cout << excel_number << endl;
  cout << endl;

  return 0;
}

void excel_count(float score){
  if(score > 86){
    excel_number++;
  }
}

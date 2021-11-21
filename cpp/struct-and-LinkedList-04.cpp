#include <iostream>
#include <iomanip>

using namespace std;
// https://www.bilibili.com/video/BV1bs41197KN?p=98

struct SpaceTraveler{
  int id;
  char name[30];
};

SpaceTraveler getNewSpaceTraveler(){
  SpaceTraveler st = {
    9527,
    "编号89757号机器人"
  };

  return st;
}

// 结构体变量作为函数的返回值
int main(){
  SpaceTraveler stupid_face = getNewSpaceTraveler();
  cout << setw(10) << stupid_face.id << " " << stupid_face.name << endl;

  cout << endl;
  return 0;
}

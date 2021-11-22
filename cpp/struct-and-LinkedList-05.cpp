#include <iostream>
#include <iomanip>

using namespace std;
// https://www.bilibili.com/video/BV1bs41197KN?p=99

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

// 指向结构体的指针
int main(){
  SpaceTraveler stupid_face = getNewSpaceTraveler();
  cout << setw(10) << stupid_face.id << " " << stupid_face.name << endl;
  cout << endl;

  SpaceTraveler *one = &stupid_face;
  cout << setw(10) << (*one).id << " " << (*one).name << endl;

  cout << endl;
  return 0;
}

#include <iostream>

using namespace std;
// https://www.bilibili.com/video/BV1bs41197KN?p=99

/*
    结构体数据类型的特性 与 普通数据类型的特性 是一致的
*/

struct SpaceTraveler{
  int id;
  char name[20];
};

// 结构体数组
int main(){
  SpaceTraveler ricks[3] = {
    137, "Rick SancheX",
    138, "Rick SancheY",
    139, "Rick SancheZ",
  };

  SpaceTraveler *one = ricks;

  cout << one->id << " " << one->name << endl;
  one++;

  cout << one->id << " " << one->name << endl;
  one++;

  cout << one->id << " " << one->name << endl;

  cout << endl;
  return 0;
}

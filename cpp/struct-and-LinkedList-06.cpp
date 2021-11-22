#include <iostream>
#include <iomanip>

using namespace std;
// https://www.bilibili.com/video/BV1bs41197KN?p=99
// struct-and-LinkedList-03.cpp

struct SpaceTraveler{
  int id;
  char name[20];
};

void renew(SpaceTraveler *st){
  st->id = 20210000 + st->id;
  for(int i = 0; st->name[i] != '\0'; i++){
    st->name[i] = toupper(st->name[i]);
  }

  cout << setw(10) << st->id << " " << st->name << endl;
}

// 结构体变量作为函数的参数, 相当于 Copy一份给函数, 和数组名作为函数参数是不同的
int main(){
  SpaceTraveler rick = {137, "Rick Sanchez"};

  renew(&rick);
  cout << setw(10) << rick.id << " " << rick.name << endl;

  cout << endl;
  return 0;
}

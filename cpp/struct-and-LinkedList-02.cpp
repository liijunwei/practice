#include <iostream>
#include <iomanip>

using namespace std;
// https://www.bilibili.com/video/BV1bs41197KN?p=98

struct SpaceTraveler{
  int id;
  char name[20];
};

int main(){
  SpaceTraveler rick1 = {137, "Rick Sanchez"};

  SpaceTraveler rick2;
  rick2 = rick1;

  rick1.id = 20210000 + rick2.id;

  for(int i = 0; rick2.name[i] != '\0'; i++){
    rick2.name[i] = toupper(rick2.name[i]);
  }

  cout << setw(10) << rick1.id << " " << rick1.name << endl;
  cout << setw(10) << rick2.id << " " << rick2.name << endl;

  cout << endl;
  return 0;
}

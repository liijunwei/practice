#include <iostream>

using namespace std;
// https://www.bilibili.com/video/BV1bs41197KN?p=97

struct student{
  char id[20];
  char name[20];
};

int main(){
  cout << "Regular lijunwei." << endl;
  student lijunwei = {"1205", {"lijunwei"}};
  cout << "id:   " << lijunwei.id << endl;
  cout << "name: " << lijunwei.name << endl;
  cout << endl;

  cout << "Regular Rick Sanchez." << endl;
  student rick = {"C137-R", "Rick Sanchez"};
  cout << "id:   " << rick.id << endl;
  cout << "name: " << rick.name << endl;
  cout << endl;

  for(int i = 0; rick.name[i] != '\0'; i++){
    rick.name[i] = toupper(rick.name[i]);
  }

  cout << "Uppercase Rick Sanchez." << endl;
  cout << "id:   " << rick.id << endl;
  cout << "name: " << rick.name << endl;
  cout << endl;

  cout << "Regular Morty." << endl;
  student morty = {"C137-M", {'M', 'o', 'r', 't', 'y', '\0'}};
  cout << "id:   " << morty.id << endl;
  cout << "name: " << morty.name << endl;
  cout << endl;

  for(int i = 0; morty.name[i] != '\0'; i++){
    morty.name[i] = toupper(morty.name[i]);
  }

  cout << "Uppercase Morty." << endl;
  cout << "id:   " << morty.id << endl;
  cout << "name: " << morty.name << endl;
  cout << endl;

  cout << endl;
  return 0;
}

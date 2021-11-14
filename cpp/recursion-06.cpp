#include <iostream>

using namespace std;

/*
汉诺塔
move(n, a, b, c)     // move n   disks from a to c via b    (target)
  move(n-1, a, c, b) // move n-1 disks from a to b via c    (step 1)
                     // move 1   disk  from a to c directly (step 2)
  move(n-1, b, a, c) // move n-1 disks from b to c via a    (step 3)

  repeat.

*/

// 将 N 个盘子, 从 A 经过 B 移动到 C
void move(int n, char from, char to, char via);

int main(){
  cout << "1 个盘子: " << endl;
  move(1, 'A', 'C', 'B');
  cout << endl;
  cout << "2 个盘子: " << endl;
  move(2, 'A', 'C', 'B');
  cout << endl;
  cout << "3 个盘子: " << endl;
  move(3, 'A', 'C', 'B');
  cout << endl;
  return 0;
}

void move(int n, char from, char to, char via){
  if(n == 1){
    cout << "把一个盘子从 " << from << " 移动到 " << to << endl;
  }
  else{
    move(n - 1, from, via, to);
    cout << "把一个盘子从 " << from << " 移动到 " << to << endl;
    move(n - 1, via, from, to);
  }

}



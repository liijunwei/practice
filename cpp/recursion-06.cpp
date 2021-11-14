#include <iostream>

using namespace std;

/*
Tower of Hanoi

move(n, a, b, c)     // move n   disks from a to c via b    (target)
  move(n-1, a, c, b) // move n-1 disks from a to b via c    (step 1)
                     // move 1   disk  from a to c directly (step 2)
  move(n-1, b, a, c) // move n-1 disks from b to c via a    (step 3)

  repeat.
*/

// Move N disks from rod-A via rod-B to rod-C
void move(int n, char from, char via, char to);

int main(){
  int n;

  cout << "Please enter number of disks: ";
  cin >> n;
  cout << "Moving " << n << " disks on 3 rods, steps as bellow: " << endl;
  move(n, 'A', 'B', 'C');

  cout << endl;
  return 0;
}

void move(int n, char from, char via, char to){ /* 最终目标: 将 n个盘子 从 from 经过 via 移动到 to */
  if(n == 1){
    cout << "Move disk from " << from << " to " << to << endl;
  }
  else{
    move(n - 1, from, to, via); /* 将 (n-1)个盘子 从 from 经过 to 移动到 via */
    cout << "Move disk from " << from << " to " << to << endl;
    move(n - 1, via, from, to); /* 将 (n-1)个盘子 从 via 经过 from 移动到 to */
  }

}



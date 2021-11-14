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

void move(int n, char from, char via, char to){ /* Ultimate goal: Move n disks from 'A` via 'B' to 'C */
  if(n == 1){
    cout << "Move disk from " << from << " to " << to << endl;
  }
  else{
    move(n - 1, from, to, via); /* Move (n-1) disks from 'A' via 'C' to 'B' */
    cout << "Move disk from " << from << " to " << to << endl;
    move(n - 1, via, from, to); /* Move (n-1) disks from 'B' via 'A' to 'C' */
  }

}



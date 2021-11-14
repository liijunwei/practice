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

void move(int n, char x, char y, char z);

int main(){

  return 0;
}

void move(int n, char x, char y, char z){


}



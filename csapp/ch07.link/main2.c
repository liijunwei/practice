#include <stdio.h>
#include "vector.h"
// #include <vector>
// #include <vector.h>

int x[2] = {1, 2};
int y[2] = {3, 4};
int z[2];

/**
sudo yum --setopt=group_package_types=mandatory,default,optional groupinstall "Development Tools" -y
sudo yum install vector.x86_64
yum provides libstdc++
sudo yum install libstdc++ -y
sudo yum update libstdc++

先跳过了
*/
int main(int argc, char const *argv[]) {
  addvec(x, y, z, 2);
  printf("z = [%d %d]\n", z[0], z[1]);

  return 0;
}

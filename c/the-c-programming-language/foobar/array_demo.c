#include <stdio.h>

// https://www.youtube.com/watch?v=90gFFdzuZMw&list=PLlcnQQJK8SUg4C8syQ63d6o44dzt_vFls&index=4

int main(int argc, char const *argv[])
{
  char *name = "LIJUNWEI";

  // 因为 name[0] <=> *(name + 0) <=> *(0 + name)
  // 所以 name[0] <=> 0[name]
  // so werid 哈哈哈哈

  printf("char is [%c]\n", name[0]);
  printf("char is [%c]\n", *(name + 0));
  printf("char is [%c]\n", 0[name]);

  printf("char is [%c]\n", name[2]);
  printf("char is [%c]\n", *(name + 2));
  printf("char is [%c]\n", 2[name]);

  return 0;
}
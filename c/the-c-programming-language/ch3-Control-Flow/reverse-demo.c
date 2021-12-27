#include <stdio.h>

#define ARR_SIZE 10

// page 51

void reverse(char s[]){

}

int main(int argc, char const *argv[])
{
  char str[ARR_SIZE] = {'b', 'x', 'z', 'y', 'x', 'y', 'z', '\0'};

  printf("before: %s\n", str);
  reverse(str);
  printf("after : %s\n", str);

  return 0;
}

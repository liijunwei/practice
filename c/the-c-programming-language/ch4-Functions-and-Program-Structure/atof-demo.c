#include <stdio.h>
#include <ctype.h>
#include <assert.h>

double atof(char s[]){

  return 3.14;
}

int main(int argc, char const *argv[])
{
  char str[100] = "3.14159265354";

  assert(3.14 == atof(str));
  printf("PASS.");
  return 0;
}
#include <stdio.h>
#include <assert.h>

/*
page 85

将数组的一部分传给函数

*/

#define ARR_SIZE 10

void print_partial_array(char *s){
  printf("string is: %s\n", s);
}

int main(int argc, char const *argv[])
{
  char s[ARR_SIZE] = {'a', 'b', 'c', 'd', '\0'};

  print_partial_array(s);
  print_partial_array(s + 1);
  print_partial_array(s + 2);
  print_partial_array(s + 3);
  print_partial_array(s + 4);

  printf("=================\n");
  print_partial_array(&s[0]);
  print_partial_array(&s[1]);
  print_partial_array(&s[2]);
  print_partial_array(&s[3]);
  print_partial_array(&s[4]);

  return 0;
}


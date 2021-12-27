#include <stdio.h>

#define ARR_SIZE 10

void shellsort(int v[], int n){

}

void print_arr(char *info, const int v[]){
  printf("%s: ", info);
  for(int i = 0; i < ARR_SIZE; i++){
    printf("%d ", v[i]);
  }
  printf("\n");
}

int main(int argc, char const *argv[])
{
  int a[ARR_SIZE] = {2022, 9, 2, 2021, 1, 3, 4, 8, 19, 1999};
  print_arr("before", a);

  shellsort(a, ARR_SIZE);
  print_arr("after ", a);

  return 0;
}

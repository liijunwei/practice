#include <stdio.h>

#define ARR_SIZE 10

void shellsort(int v[], int n){
  int gap;
  int i;
  int j;
  int temp;

  for(gap = n / 2; gap > 0; gap /= 2){
    for(i = gap; i < n; i++){
      for(j = i - gap; j >= 0 && v[j] > v[j+gap]; j -= gap){
        temp = v[j];
        v[j] = v[j+gap];
        v[j+gap] = temp;
      }
    }
  }
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

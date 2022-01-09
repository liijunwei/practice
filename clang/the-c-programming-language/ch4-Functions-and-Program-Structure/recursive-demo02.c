#include <stdio.h>

/*
page 74

快速排序算法
*/

void swap(int v[], int i, int j){
  int temp = v[i];
  v[i] = v[j];
  v[j] = temp;
}

void qsort(int v[], int left, int right){
  int i;
  int last;

  if(left >= right){ // 数组中包含的元素数小于2
    return;
  }

  swap(v, left, (left + right) / 2);
  last = left;

  for(i = left + 1; i <= right; i++){
    if(v[i] < v[left]){
      swap(v, ++last, i);
    }
  }

  swap(v, left, last);
  qsort(v, left, last - 1);
  qsort(v, left + 1, right);
}


void printarr(int v[], arr_size){
  for(int i = 0; i < arr_size; i++){
    printf("%d ", v[i]);
  }

  printf("\n");
}

#define ARR_SIZE 5

int main(int argc, char const *argv[])
{
  int a[ARR_SIZE] = {9, 1, 2, 4, 3};
  printarr(a, ARR_SIZE);

  qsort(a, 0, 4);
  printarr(a, ARR_SIZE);

  return 0;
}

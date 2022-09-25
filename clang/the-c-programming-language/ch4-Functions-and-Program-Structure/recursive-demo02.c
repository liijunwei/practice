#include <stdio.h>

/*
page 74

快速排序算法:

对于一个给定的数组, 从中选择一个元素, 以该元素为界将其余元素划分为两个自己,
一个子集中的所有元素 都小于该元素, 另一个子集的所有元素都大于或等于该元素;
对这样两个子集递归执行这一过程, 当某个自己中的元素小于2时,
这个子集就不需要再次排序, 终止递归
*/

void swap(int v[], int i, int j) {
  int temp = v[i];
  v[i] = v[j];
  v[j] = temp;
}

void qsort(int v[], int left, int right) {
  int i;
  int last;

  if (left >= right) { // 数组中包含的元素数小于2
    return;
  }

  swap(v, left, (left + right) / 2); // 划分子集
  last = left;

  for (i = left + 1; i <= right; i++) {
    if (v[i] < v[left]) {
      swap(v, ++last, i);
    }
  }

  swap(v, left, last);
  qsort(v, left, last - 1);
  qsort(v, left + 1, right);
}

void printarr(int v[], int arr_size) {
  for (int i = 0; i < arr_size; i++) {
    printf("%d ", v[i]);
  }

  printf("\n");
}

#define ARR_SIZE 5

int main(int argc, char const *argv[]) {
  int a[ARR_SIZE] = {9, 1, 2, 4, 3};
  printarr(a, ARR_SIZE);

  qsort(a, 0, ARR_SIZE - 1);
  printarr(a, ARR_SIZE);

  return 0;
}

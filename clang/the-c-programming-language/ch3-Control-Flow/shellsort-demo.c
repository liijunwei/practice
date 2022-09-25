#include <stdio.h>

#define ARR_SIZE 10

// page 51

// TODO 看明白这个排序算法

// 最外层for循环 控制两个被比较元素之间的距离, 从n/2开始, 逐步进行分析,
// 直到距离为0 中间层for循环 用于在元素间移动位置 最内层for循环
// 用于比较各各对相距gap个位置的元素, 当这两个元素逆序时, 把他们互换过来
// 由于gap的值最终将递减到1, 因此所有元素最终都会位于正确的排序位置上

void shellsort(int v[], int n) {
  int gap;
  int i;
  int j;
  int temp;

  for (gap = n / 2; gap > 0; gap /= 2) {
    for (i = gap; i < n; i++) {
      for (j = i - gap; j >= 0 && v[j] > v[j + gap]; j -= gap) {
        temp = v[j];
        v[j] = v[j + gap];
        v[j + gap] = temp;
      }
    }
  }
}

void print_arr(char *info, const int v[]) {
  printf("%s: ", info);
  for (int i = 0; i < ARR_SIZE; i++) {
    printf("%d ", v[i]);
  }
  printf("\n");
}

int main(int argc, char const *argv[]) {
  int a[ARR_SIZE] = {2022, 9, 2, 2021, 1, 3, 4, 8, 19, 1999};
  print_arr("before", a);

  shellsort(a, ARR_SIZE);
  print_arr("after ", a);

  return 0;
}

#include <stdio.h>

// page 58
float sum_elements(float a[], unsigned length) {
  float result = 0;

  for (int i = 0; i <= length; i++) {
    result += a[i];
  }

  return result;
}

// where is the bug?

int main(int argc, char const *argv[]) {

  float arr[] = {};

  unsigned length = 0;

  sum_elements(arr, length);

  return 0;
}

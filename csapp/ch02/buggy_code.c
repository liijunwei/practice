#include <stdio.h>

float sum_elements(float a[], unsigned length) {
  float result = 0;

  for (int i = 0; i <= length; i++) {
    result += a[i];
  }

  return result;
}

int main(int argc, char const *argv[])
{

  float arr[0] = {};

  unsigned length = 0;

  sum_elements(arr, length);

  return 0;
}


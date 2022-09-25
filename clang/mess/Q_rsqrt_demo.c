#include <stdio.h>

// https://www.youtube.com/watch?v=p8u_k2LIZyo
// https://en.wikipedia.org/wiki/Fast_inverse_square_root

float Q_rsqrt(float number) {
  long i;
  float x2, y;
  const float threehalfs = 1.5F;

  x2 = number * 0.5F;
  y = number;
  i = *(long *)&y;           // evil floating point bit level hacking
  i = 0x5f3759df - (i >> 1); // what the fuck?
  y = *(float *)&i;
  y = y * (threehalfs - (x2 * y * y)); // 1st iteration

  return y;
}

int main() {
  float result = 1.0 / Q_rsqrt(2);
  printf("%f", result);
}
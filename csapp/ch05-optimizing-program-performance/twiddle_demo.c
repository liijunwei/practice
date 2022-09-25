#include <stdio.h>

void twiddle1(long *xp, long *yp) {
  *xp += *yp;
  *xp += *yp;
}

void twiddle2(long *xp, long *yp) { *xp += 2 * *yp; }

/**
 * p343
 * 乍一看这两个函数好像等价; 并且 twiddle2 的效率要高一些(内存引用的次数更少)
 *
 * 但是 如果考虑 xp == yp 的情况, 就会发现情况发生了变化
 * twiddle1 会将 xp指向的值 变为原来的 4倍
 * twiddle2 会将 xp指向的值 变为原来的 3倍
 *
 * 这种两个指针可能指向同一个内存位置的情况称为内存别名使用(memory aliasing)
 */

int main(int argc, char const *argv[]) {
  long a = 100;
  printf("twiddle1: %ld\n", a);

  twiddle1(&a, &a);
  printf("twiddle1: %ld\n", a);

  printf("\n");

  long b = 100;
  printf("twiddle2 %ld\n", b);

  twiddle2(&b, &b);
  printf("twiddle2 %ld\n", b);

  return 0;
}

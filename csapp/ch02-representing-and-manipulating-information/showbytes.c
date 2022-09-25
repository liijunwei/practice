#include <stdio.h>
#include <string.h>

typedef unsigned char *byte_pointer;

void show_bytes(byte_pointer start, size_t len) {
  size_t i;

  for (i = 0; i < len; i++) {
    printf(" %.2x", start[i]);
  }

  printf("\n");
}

void show_int(int x) { show_bytes((byte_pointer)&x, sizeof(int)); }

void show_float(float x) { show_bytes((byte_pointer)&x, sizeof(float)); }

void show_pointer(void *x) { show_bytes((byte_pointer)&x, sizeof(void *)); }

void test_show_bytes(int val) {
  int ival = val;
  float fval = (float)ival;
  int *pval = &ival;

  show_int(ival);
  show_float(fval);
  show_pointer(pval);
  printf("\n");
}

// gcc csapp/ch02/showbytes.c && ./a.out
int main(int argc, char const *argv[]) {
  test_show_bytes(12345);

  show_int(3510593);
  show_float(3510593.0);

  const char *s1 = "12345";
  show_bytes((byte_pointer)s1, strlen(s1));

  const char *s2 = "abcdef";
  show_bytes((byte_pointer)s2, strlen(s2));

  return 0;
}

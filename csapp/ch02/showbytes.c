#include <stdio.h>

typedef unsigned char *byte_pointer;

void show_bytes(byte_pointer start, size_t len) {
  size_t i;

  for (i = 0; i < len; i++) {
    printf(" %.2x", start[i]);
  }

  printf("\n");
}

void show_int(int x) {
  show_bytes((byte_pointer) &x, sizeof(int));
}

void show_float(float x) {
  show_bytes((byte_pointer) &x, sizeof(float));
}

void show_float(float x) {
  show_bytes((byte_pointer) &x, sizeof(float));
}

// gcc csapp/ch02/showbytes.c && ./a.out
int main(int argc, char const *argv[]) {
  show_int(1);
  show_int(2);
  show_int(3);

  return 0;
}

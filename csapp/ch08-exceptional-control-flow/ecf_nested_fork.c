#include "csapp.h"

int main(int argc, char const *argv[]) {
  Fork();
  Fork();
  Fork();
  printf("hello\n");
  exit(0);
}

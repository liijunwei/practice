// This code is buggy
// 12.5

#include "csapp.h"

void *thread(void *vargp);

// golbal shared variable
volatile long cnt = 0; // counter

int main(int argc, char const *argv[]) {
  long niters; // number of iterations
  pthread_t tid1;
  pthread_t tid2;

  if(argc != 2) {
    printf("usage: %s <niters>\n", argv[0]);
    exit(0);
  }

  exit(0);
}

// thread routine
void *thread(void *vargp) {
  long i, niters = *((long *)vargp); // unclear...

  for(i = 0; i < niters; i++) {
    cnt++;
  }

  return NULL;
}

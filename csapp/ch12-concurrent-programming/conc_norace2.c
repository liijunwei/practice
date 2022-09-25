#include "csapp.h"

#define N 4

void *thread(void *vargp);

int main(int argc, char const *argv[]) {
  pthread_t tid[N];
  int i;

  for (i = 0; i < N; i++) {
    Pthread_create(&tid[i], NULL, thread, i);
  }

  for (i = 0; i < N; i++) {
    Pthread_join(tid[i], NULL);
  }

  exit(0);
}

// unclear: 什么叫"它假设指针至少和int一样大?"
void *thread(void *vargp) {
  int myid = (int)vargp;
  printf("Hello from thread %d\n", myid);

  return NULL;
}

#include "csapp.h"

#define N 4

void *thread(void *vargp);

int main(int argc, char const *argv[]) {
  pthread_t tid[N];
  int i;

  // `i++` 和 `int myid = *((int *)vargp);` 可能产生竞争
  // 因为传给线程的是i的地址...
  for (i = 0; i < N; i++) {
    Pthread_create(&tid[i], NULL, thread, &i);
  }

  for (i = 0; i < N; i++) {
    Pthread_join(tid[i], NULL);
  }

  exit(0);
}

void *thread(void *vargp) {
  int myid = *((int *)vargp);
  printf("Hello from thread %d\n", myid);

  return NULL;
}

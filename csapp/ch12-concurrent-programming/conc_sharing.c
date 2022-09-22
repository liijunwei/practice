// 12.4.1
#include "csapp.h"

#define N 10
void *thread(void *vargp);

char **ptr; // global variable

int main(int argc, char const *argv[]) {
  int i;

  pthread_t tid;

  char *msgs[N] = {
    "hello0",
    "hello1",
    "hello2",
    "hello3",
    "hello4",
    "hello5",
    "hello6",
    "hello7",
    "hello8",
    "hello9",
  };

  ptr = msgs;

  for(i = 0; i < N; i++) {
    Pthread_create(&tid, NULL, thread, (void *)i);
  }

  Pthread_exit(NULL);
}

// thread routine
void *thread(void *vargp) {
  int myid = (int)vargp; // unclear...

  static int cnt = 0;

  printf("[%d]: %s (cnt=%d)\n", myid, ptr[myid], ++cnt);

  return NULL;
}

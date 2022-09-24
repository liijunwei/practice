// fixed ./conc_badcnt.c

#include "csapp.h"

void *thread(void *vargp);


// golbal shared variable
volatile long cnt = 0; // counter
sem_t mutex; // semaphore that protects counter page704

int main(int argc, char const *argv[]) {
  long niters; // number of iterations
  pthread_t tid1;
  pthread_t tid2;

  if(argc != 2) {
    printf("usage: %s <niters>\n", argv[0]);
    exit(0);
  }

  niters = atoi(argv[1]);

  // create threads and wait for them to finish
  Pthread_create(&tid1, NULL, thread, &niters);
  Pthread_create(&tid2, NULL, thread, &niters);

  Pthread_join(tid1, NULL);
  Pthread_join(tid2, NULL);

  // check result
  if(cnt != (2 * niters)) {
    printf("BOOM! cnt=%ld\n", cnt);
  } else {
    printf("OK cnt=%ld\n", cnt);
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

#include "csapp.h"

#define MAXTHREADS 32

void *sum_mutex(void *vargp); // thread routine

long gsum = 0;
long nelems_per_thread;
sem_t mutex;

int main(int argc, char const *argv[]) {
  long i;
  long nelems;
  long log_nelems;
  long nthreads;
  long myid[MAXTHREADS];

  pthread_t tid[MAXTHREADS];

  if(argc != 3) {
    printf("Usage: %s <nthreads> <log_nelems>\n", argv[0]);
    exit(0);
  }

  nthreads = atoi(argv[1]);
  log_nelems = atoi(argv[2]);
  nelems = (1L << log_nelems);
  nelems_per_thread = nelems / nthreads;

  Sem_init(&mutex, 0, 1);

  for(i = 0; i < nthreads; i++) {
    myid[i] = i;
    Pthread_create(&tid[i], NULL, sum_mutex, &myid[i]);
  }

  for(i = 0; i < nthreads; i++) {
    Pthread_join(tid[i], NULL);
  }

  long expected = (nelems / 2) * (nelems - 1);

  // check the result
  if(gsum != expected) {
    printf("Error: expected=%ld actual=%ld\n", expected, gsum);
  } else {
    printf("Looking good");
  }

  exit(0);
}

void *sum_mutex(void *vargp) {
  long myid = *((long *)vargp);          // extract thread id
  long start = myid * nelems_per_thread; // start element index
  long end = start + nelems_per_thread;  // end element index
  long i;

  for(i = start; i < end; i++) {
    P(&mutex);
    gsum += i;
    V(&mutex);
  }

  return NULL;
}


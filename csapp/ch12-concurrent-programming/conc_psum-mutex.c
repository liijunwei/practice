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



  exit(0);
}


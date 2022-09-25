#include "csapp.h"

#define MAXTHREADS 32

void *sum_array(void *vargp); // thread routine

long gsum = 0;
long nelems_per_thread;
sem_t mutex;
long psum[MAXTHREADS];

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

  for(i = 0; i < nthreads; i++) {
    myid[i] = i;
    Pthread_create(&tid[i], NULL, sum_array, &myid[i]);
  }

  for(i = 0; i < nthreads; i++) {
    Pthread_join(tid[i], NULL);
  }

  for(i = 0; i < nthreads; i++) {
    gsum += psum[i];
  }

  // check the result
  if(gsum != (nelems * (nelems - 1) / 2)) {
    printf("Error: result=%ld\n", gsum);
  }

  exit(0);
}

void *sum_array(void *vargp) {
  long myid = *((long *)vargp);          // extract thread id
  long start = myid * nelems_per_thread; // start element index
  long end = start + nelems_per_thread;  // end element index
  long i;

  for(i = start; i < end; i++) {
    psum[myid] += i;
  }

  return NULL;
}


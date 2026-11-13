#include <assert.h>
#include <pthread.h>
#include <stdatomic.h>
#include <stdio.h>

enum { N = 1 };

static atomic_int sum = 0;

static void *T_sum(void *arg) {
  (void)arg;
  for (int i = 0; i < N; i++) {
    int t = atomic_load(&sum);
    t++;
    atomic_store(&sum, t);
  }
  return NULL;
}

int main(void) {
  pthread_t t1;
  pthread_t t2;

  atomic_store(&sum, 0);
  assert(pthread_create(&t1, NULL, T_sum, NULL) == 0);
  assert(pthread_create(&t2, NULL, T_sum, NULL) == 0);
  assert(pthread_join(t1, NULL) == 0);
  assert(pthread_join(t2, NULL) == 0);

  int result = atomic_load(&sum);
  printf("expected=2 got=%d\n", result);
  return result >= 1 && result <= 2 ? 0 : 1;
}

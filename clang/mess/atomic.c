#include <stdio.h>
#include <stdlib.h>
#include <pthread.h>
#include <stdatomic.h>

_Atomic int counter = 0;

void *worker_add(void *arg) {
    for (int i = 0; i < 100000; i++) {
        atomic_fetch_add(&counter, 1);
    }
    return NULL;
}

int main() {
    enum { N = 4 };
    pthread_t threads[N];

    counter = 0;
    for (int i = 0; i < N; i++)
        pthread_create(&threads[i], NULL, worker_add, NULL);
    for (int i = 0; i < N; i++)
        pthread_join(threads[i], NULL);
    printf("expected=%d  got=%d\n\n", N * 100000, counter);
}

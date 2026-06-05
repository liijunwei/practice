#include <stdio.h>
#include <stdlib.h>
#include <pthread.h>

int counter = 0;
pthread_mutex_t mtx = PTHREAD_MUTEX_INITIALIZER;

typedef struct {
    int count;
} args_t;

void *worker_mutex(void *arg) {
    int count = ((args_t *)arg)->count;
    for (int i = 0; i < count; i++) {
        pthread_mutex_lock(&mtx);
        counter++;
        pthread_mutex_unlock(&mtx);
    }
    return NULL;
}

int main(int argc, char **argv) {
    enum { N = 4 };
    int count = argc > 1 ? atoi(argv[1]) : 10000000;
    args_t a = {.count = count};
    pthread_t threads[N];

    counter = 0;
    for (int i = 0; i < N; i++)
        pthread_create(&threads[i], NULL, worker_mutex, &a);
    for (int i = 0; i < N; i++)
        pthread_join(threads[i], NULL);
    printf("expected=%d  got=%d  per_thread=%d\n", N * count, counter, count);
}

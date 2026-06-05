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

// Spinlock using atomic_flag
atomic_flag lock = ATOMIC_FLAG_INIT;
int shared = 0;

void spin_lock(atomic_flag *lk) {
    while (atomic_flag_test_and_set_explicit(lk, memory_order_acquire))
        ; // spin
}

void spin_unlock(atomic_flag *lk) {
    atomic_flag_clear_explicit(lk, memory_order_release);
}

void *worker_spinlock(void *arg) {
    for (int i = 0; i < 100000; i++) {
        spin_lock(&lock);
        shared++;           // non-atomic write, protected by spinlock
        spin_unlock(&lock);
    }
    return NULL;
}

int main() {
    enum { N = 4 };
    pthread_t threads[N];

    // --- Test 1: atomic_fetch_add ---
    printf("=== atomic_fetch_add ===\n");
    counter = 0;
    for (int i = 0; i < N; i++)
        pthread_create(&threads[i], NULL, worker_add, NULL);
    for (int i = 0; i < N; i++)
        pthread_join(threads[i], NULL);
    printf("expected=%d  got=%d\n\n", N * 100000, counter);

    // --- Test 2: spinlock (compare) ---
    printf("=== spinlock vs atomic ===\n");
    shared = 0;
    for (int i = 0; i < N; i++)
        pthread_create(&threads[i], NULL, worker_spinlock, NULL);
    for (int i = 0; i < N; i++)
        pthread_join(threads[i], NULL);
    printf("spinlock shared=%d (expected %d)\n\n", shared, N * 100000);

    // --- Test 3: compare_exchange ---
    printf("=== compare_exchange_strong ===\n");
    _Atomic int val = 10;
    int expected = 10;
    // if val == expected, set val=20 and return true
    if (atomic_compare_exchange_strong(&val, &expected, 20))
        printf("CAS succeeded: val=%d, expected=%d\n", val, expected);
    else
        printf("CAS failed: val=%d, expected=%d\n", val, expected);

    expected = 99;
    // val is 20, not 99, so CAS fails and expected is updated to val
    if (atomic_compare_exchange_strong(&val, &expected, 30))
        printf("CAS succeeded: val=%d, expected=%d\n", val, expected);
    else
        printf("CAS failed: val=%d, expected=%d (expected updated to val)\n", val, expected);

    // --- Test 4: atomic_exchange ---
    printf("\n=== atomic_exchange ===\n");
    _Atomic int x = 5;
    int old = atomic_exchange(&x, 42);
    printf("old=%d  x=%d\n", old, x);

    printf("\nAll tests done.\n");
    return 0;
}
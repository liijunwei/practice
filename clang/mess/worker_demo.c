/*
 * worker_demo.c — LCS (Longest Common Subsequence) with Producer-Consumer parallelism
 *
 * Core idea (inspired by JYY's OS course at Bilibili):
 *   1. LCS DP table → a DAG (each cell depends on ↑ ← ↖ neighbors)
 *   2. Anti-diagonals (i+j = constant) → topological levels
 *      Cells in the same anti-diagonal share NO dependencies — fully parallel.
 *   3. Producer pushes tasks level-by-level; consumer workers grab & compute.
 *
 * Build:  cc -O2 -pthread -o a.out worker_demo.c
 * Run:    ./a.out
 */

#include <pthread.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <time.h>

/* ------------------------------------------------------------------ */
/*  Data structures                                                    */
/* ------------------------------------------------------------------ */

typedef struct {
    int i, j;   /* position in the DP table */
} Task;

typedef struct {
    /* Problem dimensions & strings */
    int   m, n;
    char *X, *Y;

    /* DP table & direction marks for backtracking */
    int  **dp;      /* dp[i][j] = LCS length up to X[0..i-1], Y[0..j-1] */
    int  **dir;     /* 0=↑, 1=←, 2=↖  (encoded direction) */

    /* ---- Producer-Consumer machinery ---- */
    Task *queue;
    int   q_cap;
    int   q_head;
    int   q_tail;

    pthread_mutex_t  mtx;
    pthread_cond_t   not_empty;   /* workers wait when queue is empty     */
    pthread_cond_t   not_full;    /* producer waits when queue is full    */
    pthread_cond_t   level_done;  /* producer waits for level completion  */

    int  level_remaining;   /* tasks still unfinished in current level    */
    int  active;            /* set to 0 after the final level             */
    int  shutdown;          /* tell workers to exit                       */
} Ctx;

/* ------------------------------------------------------------------ */
/*  Utilities                                                          */
/* ------------------------------------------------------------------ */

static inline int min(int a, int b) { return a < b ? a : b; }
static inline int max(int a, int b) { return a > b ? a : b; }

/* Allocate an m×n int matrix */
static int **mat_alloc(int m, int n) {
    int **M = malloc((size_t)m * sizeof(int *));
    M[0]  = calloc((size_t)(m * n), sizeof(int));   /* contiguous */
    for (int i = 1; i < m; i++) M[i] = M[i - 1] + n;
    return M;
}

static void mat_free(int **M) {
    free(M[0]);
    free(M);
}

/* ------------------------------------------------------------------ */
/*  Consumer  —  worker thread                                         */
/* ------------------------------------------------------------------ */

static void *worker(void *arg) {
    Ctx *ctx = (Ctx *)arg;

    for (;;) {
        pthread_mutex_lock(&ctx->mtx);

        /* Wait for a task (or shutdown) */
        while (ctx->q_head == ctx->q_tail && !ctx->shutdown)
            pthread_cond_wait(&ctx->not_empty, &ctx->mtx);

        if (ctx->shutdown) {
            pthread_mutex_unlock(&ctx->mtx);
            return NULL;
        }

        /* Dequeue */
        Task t = ctx->queue[ctx->q_head];
        ctx->q_head = (ctx->q_head + 1) % ctx->q_cap;

        pthread_cond_signal(&ctx->not_full);
        pthread_mutex_unlock(&ctx->mtx);

        /* ---- Compute dp[t.i][t.j] ---- */
        int up   = ctx->dp[t.i - 1][t.j];
        int left = ctx->dp[t.i][t.j - 1];

        if (ctx->X[t.i - 1] == ctx->Y[t.j - 1]) {
            ctx->dp[t.i][t.j]  = ctx->dp[t.i - 1][t.j - 1] + 1;
            ctx->dir[t.i][t.j] = 2;               /* ↖ diagonally */
        } else if (up >= left) {
            ctx->dp[t.i][t.j]  = up;
            ctx->dir[t.i][t.j] = 0;               /* ↑ up */
        } else {
            ctx->dp[t.i][t.j]  = left;
            ctx->dir[t.i][t.j] = 1;               /* ← left */
        }

        /* Signal completion */
        pthread_mutex_lock(&ctx->mtx);
        ctx->level_remaining--;
        if (ctx->level_remaining == 0)
            pthread_cond_signal(&ctx->level_done);
        pthread_mutex_unlock(&ctx->mtx);
    }
}

/* ------------------------------------------------------------------ */
/*  Producer  —  main thread feeds anti-diagonal levels                */
/* ------------------------------------------------------------------ */

/* All cells where  i + j == level  (1 ≤ i ≤ m, 1 ≤ j ≤ n).
 * These cells form one anti-diagonal and are independent of each other.
 * CRITICAL: level_remaining must be set BEFORE any worker can decrement it,
 * otherwise a worker finishing a task from the new level would decrement the
 * stale (0) counter from the previous level → deadlock.  We hold the mutex
 * across the entire enqueue loop so no worker sees tasks before the counter. */
static void fill_level(Ctx *ctx, int level) {
    pthread_mutex_lock(&ctx->mtx);

    for (int i = max(1, level - ctx->n); i <= min(ctx->m, level - 1); i++) {
        int j = level - i;
        if (j < 1 || j > ctx->n) continue;

        Task t = { i, j };

        while ((ctx->q_tail + 1) % ctx->q_cap == ctx->q_head)
            pthread_cond_wait(&ctx->not_full, &ctx->mtx);   /* queue full — wait */

        ctx->queue[ctx->q_tail] = t;
        ctx->q_tail = (ctx->q_tail + 1) % ctx->q_cap;

        ctx->level_remaining++;                 /* atomically count tasks */
        pthread_cond_signal(&ctx->not_empty);
    }

    pthread_mutex_unlock(&ctx->mtx);
}

/* ------------------------------------------------------------------ */
/*  Backtrack to reconstruct the LCS string                            */
/* ------------------------------------------------------------------ */

static char *backtrack(Ctx *ctx) {
    int len = ctx->dp[ctx->m][ctx->n];
    char *lcs = malloc((size_t)(len + 1));
    lcs[len] = '\0';

    int i = ctx->m, j = ctx->n, k = len - 1;
    while (i > 0 && j > 0) {
        if (ctx->dir[i][j] == 2) {        /* ↖  — match */
            lcs[k--] = ctx->X[i - 1];
            i--; j--;
        } else if (ctx->dir[i][j] == 0) { /* ↑ */
            i--;
        } else {                           /* ← */
            j--;
        }
    }
    return lcs;
}

/* ------------------------------------------------------------------ */
/*  Main                                                               */
/* ------------------------------------------------------------------ */

int main(int argc, char **argv) {
    /* Default strings; override via command line */
    char *X = "ABCBDAB";
    char *Y = "BDCABA";
    if (argc >= 3) { X = argv[1]; Y = argv[2]; }

    int m = (int)strlen(X);
    int n = (int)strlen(Y);
    int num_workers = 4;

    if (argc >= 4) num_workers = atoi(argv[3]);
    if (num_workers < 1) num_workers = 1;

    /* ---- Allocate context ---- */
    Ctx ctx = {
        .m = m, .n = n, .X = X, .Y = Y,
        .dp  = mat_alloc(m + 1, n + 1),
        .dir = mat_alloc(m + 1, n + 1),
        .active = 1, .shutdown = 0,
    };

    /* Queue: worst case — longest anti-diagonal is min(m,n) cells */
    ctx.q_cap  = min(m, n) * 2 + 4;
    ctx.queue  = malloc((size_t)ctx.q_cap * sizeof(Task));
    ctx.q_head = ctx.q_tail = 0;

    pthread_mutex_init(&ctx.mtx, NULL);
    pthread_cond_init(&ctx.not_empty,  NULL);
    pthread_cond_init(&ctx.not_full,   NULL);
    pthread_cond_init(&ctx.level_done, NULL);

    /* ---- Spawn workers ---- */
    pthread_t *threads = malloc((size_t)num_workers * sizeof(pthread_t));
    for (int i = 0; i < num_workers; i++)
        pthread_create(&threads[i], NULL, worker, &ctx);

    /* ---- Producer loop: feed one anti-diagonal level at a time ---- */
    /* Level k = i + j goes from 2 to m+n.
     * dp[0][*] and dp[*][0] are already 0 (base cases). */

    clock_t start = clock();

    for (int level = 2; level <= m + n; level++) {
        /* Reset counter (under mutex so no stale decrements land here) */
        pthread_mutex_lock(&ctx.mtx);
        ctx.level_remaining = 0;
        pthread_mutex_unlock(&ctx.mtx);

        fill_level(&ctx, level);  /* enqueues tasks + sets level_remaining */

        /* Wait until all tasks in this level are consumed & computed */
        pthread_mutex_lock(&ctx.mtx);
        while (ctx.level_remaining > 0)
            pthread_cond_wait(&ctx.level_done, &ctx.mtx);
        pthread_mutex_unlock(&ctx.mtx);
    }

    clock_t end = clock();

    /* ---- Shut down workers ---- */
    pthread_mutex_lock(&ctx.mtx);
    ctx.shutdown = 1;
    pthread_cond_broadcast(&ctx.not_empty);
    pthread_mutex_unlock(&ctx.mtx);

    for (int i = 0; i < num_workers; i++)
        pthread_join(threads[i], NULL);

    /* ---- Results ---- */
    char *lcs = backtrack(&ctx);
    printf("X      = %s (len %d)\n", X, m);
    printf("Y      = %s (len %d)\n", Y, n);
    printf("LCS    = %s  (len %d)\n", lcs, ctx.dp[m][n]);
    printf("Time   = %.3f ms  (%d workers)\n",
           1000.0 * (end - start) / CLOCKS_PER_SEC, num_workers);

    /* ---- DAG / topological info ---- */
    printf("\nDAG topological levels (anti-diagonals):\n");
    for (int level = 2; level <= m + n; level++) {
        int cnt = 0;
        for (int i = max(1, level - n); i <= min(m, level - 1); i++) {
            int j = level - i;
            if (j >= 1 && j <= n) cnt++;
        }
        printf("  level %2d  →  %2d independent tasks\n", level, cnt);
    }

    /* ---- Cleanup ---- */
    free(lcs);
    mat_free(ctx.dp);
    mat_free(ctx.dir);
    free(ctx.queue);
    free(threads);
    pthread_mutex_destroy(&ctx.mtx);
    pthread_cond_destroy(&ctx.not_empty);
    pthread_cond_destroy(&ctx.not_full);
    pthread_cond_destroy(&ctx.level_done);

    return 0;
}

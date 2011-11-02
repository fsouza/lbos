#include <stdio.h>
#include <stdlib.h>
#include <pthread.h>

#define N_CHILD 50

pthread_mutex_t mutex = PTHREAD_MUTEX_INITIALIZER;

typedef struct shared_st {
    int counter, end;
} Shared;

Shared *make_shared(int end) {
    Shared *shared = (Shared *)malloc(sizeof(Shared));
    shared->counter = 0;
    shared->end = end;

    return shared;
}

void desalloc_shared(Shared *shared) {
    free(shared);
}

void *increment(void *arg) {
    Shared *shared = (Shared *)arg;
    printf("Starting to increment from thread (current counter)...\n");
    while (1) {
        pthread_mutex_lock(&mutex);
        if (shared->counter < shared->end) {
            shared->counter++;
            pthread_mutex_unlock(&mutex);
        } else {
            pthread_mutex_unlock(&mutex);
            break;
        }
    }

    printf("Finish incrementing from thread.\n");
    pthread_exit(NULL);
}

int main(void) {
    int i;
    Shared *shared = make_shared(1000000);
    pthread_t child[N_CHILD];

    for (i = 0; i < N_CHILD; i++) {
        pthread_create(&child[i], NULL, increment, (void *)shared);
    }

    for (i = 0; i < N_CHILD; i++) {
        pthread_join(child[i], NULL);
    }

    printf("Final value: %d.\n\n\n", shared->counter);
    desalloc_shared(shared);
    return 0;
}

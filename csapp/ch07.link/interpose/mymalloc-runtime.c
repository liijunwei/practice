#ifdef RUNTIME

#define _GNU_SOURCE
#include <stdio.h>
#include <stdlib.h>
#include <dlfcn.h>

// malloc wrapper function
void *malloc(size_t size) {
  void *

  printf("malloc(%d)=%p\n", (int)size, ptr);

  return ptr;
}

// free wrapper function
void free(void *ptr) {

  printf("frees(%p)\n", ptr);
}

#endif

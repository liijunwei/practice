#ifdef COMPILETIME

#include <malloc.h> // 标准malloc.h头文件
#include <stdio.h>

// malloc wrapper function
void *mymalloc(size_t size) {
  void *ptr = malloc(size);
  printf("malloc(%d)=%p\n", (int)size, ptr);

  return ptr;
}

// free wrapper function
void myfree(void *ptr) {
  free(ptr);
  printf("free(%p)\n", ptr);
}

#endif

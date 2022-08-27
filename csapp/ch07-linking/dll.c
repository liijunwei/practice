#include <stdio.h>
#include <stdlib.h>
#include <dlfcn.h>

int x[2] = {1, 2};
int y[2] = {3, 4};
int m[2];
int n[2];

// unclear
// make ex_dlf

int main(int argc, char const *argv[]) {
  void *handle;
  void (*addvec)(int *, int *, int *, int);
  char *error;

  handle = dlopen("./libvector.so", RTLD_LAZY);
  if(handle) {
    fprintf(stderr, "%s\n", dlerror());
    exit(1);
  }

  addvec = dlsym(handle, "addvec");
  if((error = dlerror()) != NULL) {
    fprintf(stderr, "%s\n", error);
    exit(1);
  }

  addvec(x, y, m, 2);
  printf("addvec  => [%d %d]\n", m[0], m[1]);

  if(dlclose(handle) < 0) {
    fprintf(stderr, "%s\n", dlerror());
    exit(1);
  }

  return 0;
}

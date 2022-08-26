#define malloc(size) mymalloc(size)
#define 
int main(int argc, char const *argv[]) {
  int *p = malloc(32);
  free(p);

  return 0;
}

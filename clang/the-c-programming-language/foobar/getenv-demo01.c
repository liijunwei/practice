#include <stdio.h>
#include <stdlib.h>

/* man 3 getenv */
int main(int argc, char const *argv[]) {
  char *str;
  char const *env_path = "PATH";

  str = getenv(env_path);
  printf("PATH: %s\n", str);

  return 0;
}

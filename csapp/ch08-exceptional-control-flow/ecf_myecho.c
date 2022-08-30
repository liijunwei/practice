#include "csapp.h"

int main(int argc, char const *argv[], char const *envp[]) {
  int i;

  printf("Command-Line arguments:\n");

  for(i = 0; argv[i] != NULL; i++) {
    printf("    argv[%2d]: %s\n", i, argv[i]);
  }

  printf("\n");

  printf("Environment variables:\n");
  for(i = 0; envp[i] != NULL; i++) {
    printf("    envp[%2d]: %s\n", i, envp[i]);
  }

  exit(0);
}

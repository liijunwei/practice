#include <stdio.h>
#include <errno.h>
#include <stdlib.h>

void unix_error(char *msg) {
  fprintf(stderr, "%s: %s\n", msg, strerror(errno));
  exit(0);
}

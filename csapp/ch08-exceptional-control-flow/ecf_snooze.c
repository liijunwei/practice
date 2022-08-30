#include "csapp.h"

unsigned int snooze(unsigned int secs) {
  printf("Sleeping...\n");
  unsigned int rc = sleep(secs);
  printf("Slept for %d of %d secs.\n", secs - rc, secs);

  return rc;
}

int main(int argc, char const *argv[]) {
  // question: how to interrupt?
  snooze(10);

  exit(0);
}

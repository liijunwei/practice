#include "csapp.h"

void handler(int sig) {
  return; // catch he signal and return
}

unsigned int snooze(unsigned int secs) {
  printf("Sleeping...\n");
  unsigned int rc = sleep(secs);
  printf("Slept for %d of %d secs.\n", secs - rc, secs);

  return rc;
}

int main(int argc, char const *argv[]) {
  if(argc != 2)

  return 0;
}

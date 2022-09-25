#include "csapp.h"

// capture ctrl+c signal
// print message
// stop the program

// signal safe version of `./ecf_sigint.c`
// p535
void sigint_handler(int sig) {
  sio_puts("Caught SIGINT!(with safe output)\n");
  _exit(0); // safe exit
}

int main(int argc, char const *argv[]) {
  // install the SIGINT handler
  // man signal
  if (signal(SIGINT, sigint_handler) == SIG_ERR) {
    unix_error("signal error...");
  }

  Pause(); // wait for the receipt of a signal

  exit(0);
}

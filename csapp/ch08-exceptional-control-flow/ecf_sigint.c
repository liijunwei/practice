#include "csapp.h"

// capture ctrl+c signal
// print message
// stop the program

void sigint_handler(int sig) {
  printf("Caught SIGINT!(with unsafe output)\n"); // printf不是异步信号安全的函数
                                                  // p535
  exit(0);
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

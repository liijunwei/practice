#include "csapp.h"

int main(int argc, char const *argv[]) {
  Socket(AF_INET, SOCK_STREAM, 0);                // 3
  Socket(AF_INET, SOCK_STREAM, 0);                // 4
  Socket(AF_INET, SOCK_STREAM, 0);                // 5
  int clientfd = Socket(AF_INET, SOCK_STREAM, 0); // 6

  printf("clientfd: %d\n", clientfd);

  exit(0);
}

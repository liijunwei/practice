#include "csapp.h"

int main(int argc, char const *argv[]) {
  struct in_addr inaddr; // address in network byte order
  int rc;

  if (argc != 2) {
    fprintf(stderr, "usage: %s <dotted-decimal>\n", argv[0]);
    exit(0);
  }

  rc = inet_pton(AF_INET, argv[1], &inaddr);

  if (rc == 0) {
    app_error("inet_pton error: invalid dotted-decimal address");
  } else if (rc < 0) {
    unix_error("inet_pton error");
  }

  printf("0x%x\n", ntohl(inaddr.s_addr));

  exit(0);
}

#include "csapp.h"

// hex2dd.c
int main(int argc, char const *argv[]) {
  struct in_addr inaddr; // address in network byte order
  uint32_t addr;         // address in host byte order
  char buf[MAXBUF];

  if(argc != 2) {
    fprintf(stderr, "usage: %s <hex number>\n", argv[0]);
    exit(0);
  }

  sscanf(argv[1], "%x", &addr);
  inaddr.s_addr = htonl(addr);

  if(!inet_ntop(AF_INET, &inaddr, buf, MAXBUF)) {
    unix_error("inet_ntop");
  }

  printf("%s\n", buf);

  exit(0);
}

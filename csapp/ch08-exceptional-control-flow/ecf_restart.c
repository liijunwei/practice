#include "csapp.h"

jmp_buf buf;

void handler(int sig) {
  siglongjmp(buf, 1);
}

int main(int argc, char const *argv[]) {
  if(!sigsetjmp(buf, 1)) {
    Signal(SIGINT, handler);
    sio_puts("starting\n");
  } else {
    sio_puts("restarting\n");
  }

  while(1) {
    Sleep(1);
    sio_puts("processing...\n");
  }

  exit(0);
}

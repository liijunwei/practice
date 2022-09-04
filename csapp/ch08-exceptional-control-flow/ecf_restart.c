#include "csapp.h"

// p549(图8-44)
// c++ 和 java中提供的异常机制是较高层次的, 是c语言的setjmp和longjmp函数的更加结构化的版本
// 你可以把 try语句中的catch子句看做类似setjmp函数, 相似的, throw语句类似于longjmp函数 (???)

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

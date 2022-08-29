#include <stdio.h>
#include <unistd.h>
#include <stdlib.h>

pid_t Fork(void) {
  pid_t pid;

  if((pid = fork()) < 0) {
    unix_error("Fork error");
  }

  return pid;
}

int main(int argc, char const *argv[]) {


  return 0;
}

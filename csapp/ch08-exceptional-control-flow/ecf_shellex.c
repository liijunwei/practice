#include "csapp.h"
#define MAXARGS 128

void eval(char *cmdline);
int parseline(char *buf, char **argv);
int builtin_command(char **argv);

int main(int argc, char const *argv[], char const *envp[]) {
  char cmdline[MAXLINE];

  while(1) {
    printf("> ");
    Fgets(cmdline, MAXLINE, stdin);

    if(feof(stdin)) {
      exit(0);
    }

    eval(cmdline);
  }

  exit(0);
}



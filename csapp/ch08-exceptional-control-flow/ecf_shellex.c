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

void eval(char *cmdline) {
  char *argv[MAXARGS];
  char buf[MAXLINE];
  int bg;
  pid_t pid;

  strcpy(buf, cmdline);
  bg = parseline(buf, argv);

  if(argv[0] == NULL) {
    return;
  }

  if(!builtin_command(argv)) {
    if((pid = Fork()) == 0) {
      if(execv(argv[0], argv, environ) < 0) {
        printf("%s: Command not found.\n", argv[0]);
        exit(0);
      }
    }

    if(!bg) {
      int status;
      if(waitpid(pid, &status, 0) < 0) {
        unix_error("waitfg: waitpid error");
      }
    } else {
      printf("%d %s", pid, cmdline);
    }
  }

  return;
}


int builtin_command(char **argv) {
  if(!strcmp(argv[0], "quit"))
}





#include "csapp.h"

int main(int argc, char const *argv[]) {
  struct stat stat;
  char *type;
  char *readok;

  Stat(argv[1], &stat);

  // check file type
  if(S_ISREG(stat.st_mode)) {
    type = "regular";
  } else if(S_ISDIR(stat.st_mode)) {
    type = "directory";
  } else {
    type = "others";
  }

  // check read access
  if((stat.st_mode & S_IRUSR)) {
    readok = "yes";
  } else {
    readok = "no";
  }

  printf("type: %s, read: %s\n", type, readok);

  exit(0);
}

/*
page 144

使用fgets很容易实现getline函数

*/

#include <stdio.h>
#include <string.h>

int custom_getline(char *line, int max) {
  if (fgets(line, max, stdin) == NULL) {
    return 0;
  } else {
    return strlen(line);
  }
}

int main(int argc, char const *argv[]) {
  char line[100];

  custom_getline(line, 100);
  printf("you entered: %s", line);

  return 0;
}

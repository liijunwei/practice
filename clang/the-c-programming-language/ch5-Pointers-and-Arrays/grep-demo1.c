#include <stdio.h>
#include <string.h>

/*
page 100

针对4-1中模式查找程序的改进
*/

#define MAXLINE 1000
int custom_getline(char s[], int max);

int main(int argc, char const *argv[]) {
  char line[MAXLINE];
  int found = 0;

  if (argc != 2) {
    printf("Usage: find pattern\n");
  } else {
    while (custom_getline(line, MAXLINE) > 0) {
      if (strstr(line, argv[1]) != NULL) {
        printf("%s", line);
        found++;
      }
    }
  }

  return found;
}

int custom_getline(char s[], int max) {
  int c;
  int i = 0;

  while (--max > 0 && (c = getchar()) != EOF && c != '\n') {
    s[i] = c;
    i++;
  }

  if (c == '\n') {
    s[i] = c;
    ++i;
  }

  s[i] = '\0';

  return i;
}

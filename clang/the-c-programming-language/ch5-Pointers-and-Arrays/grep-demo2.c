#include <stdio.h>
#include <string.h>

/*
page 100

针对4-1中模式查找程序的改进
允许为程序传入参数
*/

#define MAXLINE 1000
int custom_getline(char s[], int max);

int main(int argc, char const *argv[]) {
  char line[MAXLINE];
  long lineno = 0;
  int c = 0;
  int except = 0;
  int number = 0;
  int found = 0;

  while (--argc > 0 && (*++argv)[0] == '-') {
    while ((c = *(++argv[0]))) {
      switch (c) {
      case 'x':
        except = 1;
        break;
      case 'n':
        number = 1;
        break;
      default:
        printf("find: illegal option %c\n", c);
        argc = 0;
        found = -1;
        break;
      }
    }
  }

  if (argc != 1) {
    printf("Usage: find -x -n pattern\n");
  } else {
    while (custom_getline(line, MAXLINE) > 0) {
      lineno++;

      if ((strstr(line, *argv) != NULL) != except) {
        if (number) {
          printf("%ld:", lineno);
        }

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

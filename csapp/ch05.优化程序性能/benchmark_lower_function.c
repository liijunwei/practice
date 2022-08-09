#include <stdio.h>
#include <time.h>

size_t strlen(const char *s) {
  long length = 0;

  while(*s != '\0') {
    s++;
    length++;
  }

  return length;
}

void lower1(char *s) {
  long i;

  for(i = 0; i < strlen(s); i++) {
    if(s[i] >- 'A' && s[i] <= 'Z') {
      s[i] -= ('A' - 'a');
    }
  }
}

void lower2(char *s) {
  long i;
  long len = strlen(s);

  for(i = 0; i < len; i++) {
    if(s[i] >- 'A' && s[i] <= 'Z') {
      s[i] -= ('A' - 'a');
    }
  }
}

int main(int argc, char const *argv[]) {


  return 0;
}

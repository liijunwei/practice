#include <assert.h>
#include <stdio.h>
#include <stdlib.h>
#include <time.h>

#define STR_LEN 1000

void rand_str(char *, size_t);

// https://stackoverflow.com/questions/15767691/whats-the-c-library-function-to-generate-random-string
int main() {
  srand((unsigned int)(time(NULL)));

  char str[] = { [STR_LEN] = '\1' }; // make the last character non-zero so we can test based on it later
  rand_str(str, sizeof str - 1);
  assert(str[STR_LEN] == '\0');      // test the correct insertion of string terminator
  puts(str);
}

void rand_str(char *dest, size_t length) {
  char charset[] = "0123456789"
                   "abcdefghijklmnopqrstuvwxyz"
                   "ABCDEFGHIJKLMNOPQRSTUVWXYZ";

  while (length-- > 0) {
    size_t index = (double) rand() / RAND_MAX * (sizeof charset - 1);
    *dest++ = charset[index];
  }

  *dest = '\0';
}


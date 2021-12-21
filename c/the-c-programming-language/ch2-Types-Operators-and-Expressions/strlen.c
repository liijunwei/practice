#include <stdio.h>

// page 30

int custom_strlen(const char s[]){
  int i = 0;

  while(s[i] != '\0'){
    ++i;
  }

  return i;
}

int main(int argc, char const *argv[])
{
  char *username = "Rubyist Seamon";
  printf("username length: %d\n", custom_strlen(username));
  return 0;
}


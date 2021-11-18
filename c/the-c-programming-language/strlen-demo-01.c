#include <stdio.h>
#include <string.h>

int costom_strlen(char s[]){
  int i = 0;
  while(s[i] != '\0'){ i++; }

  return i;
}

int main(){
  char c[5] = {'a', 'b', 'c', 'd'};

  printf("length of c is %d\n", costom_strlen(c));
  printf("length of c is %lu\n", strlen(c));
}
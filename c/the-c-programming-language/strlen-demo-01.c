#include <stdio.h>

int c_strlen(char s[]){
  int i = 0;
  while(s[i] != '\0'){ i++; }

  return i;
}

int main(){
  char c[5] = {'a', 'b', 'c', 'd'};

  printf("length of c is %d\n", c_strlen(c));
}
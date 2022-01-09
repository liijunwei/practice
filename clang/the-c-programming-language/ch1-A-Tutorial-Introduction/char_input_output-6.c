#include <stdio.h>

// page 12
int main(){
  printf("please enter chars, `ctrl+d` to exit\n");
  printf("\n");

  double nc;
  for(nc = 0; getchar() != EOF; ++nc){}

  printf("char count is %.0f\n", nc);
}

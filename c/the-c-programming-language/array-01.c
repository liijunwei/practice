#include <stdio.h>

#define ARRAY_SIZE 10
// 统计各个数字, 空白符及其他字符出现的次数

int main(){
  int c;
  int i;
  int nwhite = 0;
  int nother = 0;

  int ndight[ARRAY_SIZE];

  printf("Array Initializing...\n");
  for(i = 0; i < ARRAY_SIZE; ++i){
    ndight[i] = 0;
  }
  printf("Array Initialized!\n");
  printf("please enter chars, `ctrl+d` to exit\n");

  while((c = getchar()) != EOF){
    if(c >= '0' && c <= '9'){
      ++ndight[c - '0'];
    }
    else if(c == ' ' || c == '\n' || c == '\t'){
      ++nwhite;
    }
    else{
      ++nother;
    }
  }

  printf("dights = ");

  for(i = 0; i < ARRAY_SIZE; ++i){
    printf(" %d", ndight[i]);
  }

  printf(", white space = %d, other = %d\n", nwhite, nother);

}

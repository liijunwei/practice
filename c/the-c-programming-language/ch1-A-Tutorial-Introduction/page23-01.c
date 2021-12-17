#include <stdio.h>
#include <string.h>

// page 23
#define LENTH_MAX 1000

int max; // 目前为止发现的最长的行
char line[LENTH_MAX]; // 当前的输入行
char longest[LENTH_MAX]; // 最长的行

int custom_getline(void);
void custom_copy(void);

int main(){
  int len;
  extern int max;
  extern char longest[];

  max = 0;
  while((len = custom_getline()) > 0){
    if(len > max){
      max = len;
      custom_copy();
    }
  }

  if(max > 0){ // 如果存在这样的行
    printf("the longest line is: \n");
    printf("%s", longest);
  }

  return 0;
}

int custom_getline(void){
  int c;
  int i;
  extern char line[];

  for(i = 0; i < LENTH_MAX - 1 && (c = getchar()) != EOF && c != '\n'; ++i){
    line[i] = c;
  }

  if(c == '\n'){
    line[i] = c;
    ++i;
  }

  line[i] = '\0';

  return i;
}

void custom_copy(void){
  int i = 0;
  extern char line[];
  extern char longest[];

  while((longest[i] = line[i]) != '\0'){
    ++i;
  }
}
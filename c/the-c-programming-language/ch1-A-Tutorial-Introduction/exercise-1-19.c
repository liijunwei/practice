#include <stdio.h>
#include <string.h>

// page 22

// 编写函数 reverse(s), 将字符串s中的字符顺序颠倒过来, 使用该函数编写一个程序, 每次颠倒一个输入行中的字符顺序

int custom_getline(char s[], int limit);
int custom_reverse(char s[]);
#define LENTH_MAX 1000

int main(){
  char buffer[100];

  while(custom_getline(buffer, LENTH_MAX) > 0){

  }

  return 0;
}

int custom_getline(char s[], int limit){
  int c;
  int i;

  for(i = 0; i < limit - 1 && (c = getchar()) != EOF && c != '\n'; ++i){
    s[i] = c;
  }

  if(c == '\n'){
    s[i] = c;
    ++i;
  }

  s[i] = '\0';

  return i;
}

void custom_reverse(char s[]){

}


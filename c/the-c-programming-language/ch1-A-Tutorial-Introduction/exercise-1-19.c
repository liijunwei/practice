#include <stdio.h>
#include <string.h>

// page 22

// 编写函数 reverse(s), 将字符串s中的字符顺序颠倒过来, 使用该函数编写一个程序, 每次颠倒一个输入行中的字符顺序

int custom_getline(char s[], int limit);
void custom_reverse(char s[]);
#define LENTH_MAX 1000

int main(){
  char buffer[100];

  while(custom_getline(buffer, LENTH_MAX) > 0){
    custom_reverse(buffer);
    printf("%s", buffer);
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
  int i = 0;
  int j;
  char temp;

  while(s[i] != '\0'){ // find the end of string
    ++i;
  }

  --i; // back off from '\0'

  if(s[i] == '\n'){
    --i;
  }

  j = 0;
  while(j < i){
    // 字符串内, 首尾相互交换, 首部步进, 尾部步退, 精巧
    temp = s[j]; // swap the characters
    s[j] = s[i];
    s[i] = temp;
    --i;
    ++j;
  }
}


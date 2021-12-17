#include <stdio.h>
#include <string.h>

// page 22

// 编写程序, 删除每个输入行末尾的空格及制表符, 并删除完全是空格的行

int custom_getline(char s[], int limit);
int remove_blank(char s[]);
#define LENTH_MAX 1000

int main(){
  char buffer[100];

  while(custom_getline(buffer, LENTH_MAX) > 0){
    if(remove_blank(buffer) > 0){
      printf("%s", buffer);
    }
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

int remove_blank(char s[]){
  printf("string length is %ld\n", strlen(s));
  int i = 0;
  while(s[i] != '\n'){
    ++i; // find new line character
  }

  --i; // backoff from '\n'

  while(i > 0 && (i == ' ' || i == '\t')){
    --i;
  }

  if(i > 0){ // non blank line?
    ++i;
    s[i] = '\n'; // put back '\n'
    ++i;
    s[i] = '\0'; // put back '\0'
  }

  printf("line->\n %sreturns %d\n", s, i);
  printf("string length is %ld\n", strlen(s));
  return i;
}


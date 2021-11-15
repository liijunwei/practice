#include <stdio.h>

#define MAXLINE 15

// 将一行读入s中, 并返回其长度
int custom_getline(char line[], int maxline);

// 将 from 复制到 to, 假定 to 足够大
void copy(char to[], char from[]);

int main(){
  int len;
  int max = 0;

  char line[MAXLINE];
  char longest[MAXLINE];

  while((len = custom_getline(line, MAXLINE)) > 0){
    if(len > max){
      max = len;
      copy(longest, line);
    }
  }

  if(max > 0){
    printf("%s ", longest);
  }

  return 0;
}

int custom_getline(char line[], int maxline){
  int c;
  int i;

  for(i = 0; (i < maxline && (c = getchar() != EOF) && c != '\n'); ++i){
    line[i] = c;
  }

  if(c == '\n'){
    line[i] = c;
    ++i;
  }

  line[i] = '\n';

  return i;
}

void copy(char to[], char from[]){
  int i = 0;

  while((to[i] = from[i]) != '\n'){
    ++i;
  }
}



#include <stdio.h>

// page 22

// TODO 没明白需求是什么

// 修改打印最长文本行的main函数, 使之可以打印任意长度的输入行的长度, 并尽可能多地打印文本

// 将输入的行读入s, 并返回其长度
int custom_getline(char s[], int limit);
// 字符数组复制
void custom_copy(char to[], char from[]);

#define MAXLINE 1000 // 允许的输入行的最大长度

int main(){
  int len;               /* 当前行长度 */
  int max = 0;           /* 目前为止发现的最长行的长度 */
  char line[MAXLINE];    /* 当前的输入行 */
  char longest[MAXLINE]; /* 保存最长的行 */

  while((len = custom_getline(line, MAXLINE)) > 0){
    if(len > max){
      max = len;
      printf("copying maxline.....\n");
      custom_copy(longest, line);
    }
  }

  if(max > 0){
    printf("----------longest---------\n");
    printf("%s", longest);
  }

  return 0;
}

int custom_getline(char s[], int limit){
  int c;
  int i;
  int j = 0;

  for(i = 0; (c = getchar()) != EOF && c != '\n'; ++i){
    if(i < limit - 2){
      s[j] = c; // line length still in boundary
      ++j;
    }

    if(c == '\n'){
      s[j] = c;
      j++;
      i++;
    }

    s[j] = '\0';
    return i;
  }

  return 0;
}

void custom_copy(char to[], char from[]){
  int i = 0;

  while((to[i] = from[i]) != '\0'){
    ++i;
  }
}


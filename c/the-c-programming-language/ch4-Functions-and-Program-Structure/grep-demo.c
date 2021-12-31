#include <stdio.h>

/*
page 57

Ah Love, could you an I with Fate conspire
To grasp this sorry Scheme Of Things entire,
Re-mould it nearer to the Heart's Desire!
*/

#define MAXLINE 1000 // 输入行的最大长度

// 将行保存到s中, 并返回该行的行数
int custom_getline(char s[], int max){
  int c;
  int i = 0;

  while(--max > 0 && (c = getchar()) != EOF && c != '\n'){
    s[i] = c;
    i++;
  }

  if(c == '\n'){
    s[i] = c;
    ++i;
  }

  s[i] = '\0';

  return i;
}

// 返回t在s中的位置, 若未找到则返回-1
// s -> source
// t -> searchTarget
int strindex(char s[], char t[]){
  int i;
  int j;
  int k;

  for(i = 0; s[i] != '\0'; i++){
    for(j = i, k = 0; t[k] != '\0' && s[j] == t[k]; j++, k++){
      ;
    }

    if(k > 0 && t[k] == '\0'){
      return i;
    }
  }

  return -1;
}

char pattern[] = "ould";

// 找出所有与模式匹配的行
int main(int argc, char const *argv[])
{
  char line[MAXLINE];
  int found = 0;
  int index;

  while(custom_getline(line, MAXLINE) > 0){
    if((index = strindex(line, pattern)) >= 0){
      printf("%s(index: %d)\n", line, index);
      found++;
    }
  }

  return found;
}


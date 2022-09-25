#include <stdio.h>

/*
page 60

编写函数 strrindex(s, t), 他返回字符串t在s中最右边出现的位置
如果s中不包含t, 则返回-1

*/

#define MAXLINE 1000 // 输入行的最大长度

// 将行保存到s中, 并返回该行的行数
int custom_getline(char s[], int max) {
  int c;
  int i = 0;

  while (--max > 0 && (c = getchar()) != EOF && c != '\n') {
    s[i] = c;
    i++;
  }

  if (c == '\n') {
    s[i] = c;
    ++i;
  }

  s[i] = '\0';

  return i;
}

// 返回t在s中的位置, 若未找到则返回-1
// s -> source
// t -> searchTarget
// 使用 rpos存储最后一次出现的位置
int strindex(char s[], char t[]) {
  int i;
  int j;
  int k;
  int rpos = -1;

  for (i = 0; s[i] != '\0'; i++) {
    for (j = i, k = 0; t[k] != '\0' && s[j] == t[k]; j++, k++) {
      ;
    }

    if (k > 0 && t[k] == '\0') {
      rpos = i;
    }
  }

  return rpos;
}

char pattern[] = "ould";

// 找出所有与模式匹配的行
int main(int argc, char const *argv[]) {
  char line[MAXLINE];
  int found = 0;
  int index;

  while (custom_getline(line, MAXLINE) > 0) {
    if ((index = strindex(line, pattern)) >= 0) {
      printf("%s(index: %d)\n", line, index);
      found++;
    }
  }

  return found;
}

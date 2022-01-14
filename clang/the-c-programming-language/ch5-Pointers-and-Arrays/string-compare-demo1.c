#include <stdio.h>

/*
page 91

*/

// 根据s按照字典顺序小于/等于/大于的结果分别返回负整数/0/正整数
int custom_strcmp(char *s, const char *t){
  int i;

  for(i = 0; s[i] == t[i]; i++){
    if(s[i] == '\0'){
      return 0;
    }
  }

  return s[i] - t[i];
}

void format_and_print(char *s, char *t){
  printf("[%s] |\t [%s] -> %d\n", s, t, custom_strcmp(s, t));
}

int main(int argc, char const *argv[])
{
  format_and_print("abc", "abd");
  format_and_print("abd", "abc");
  format_and_print("abc", "abc");
  format_and_print("abc", "acc");
  format_and_print("abc", "abcd");

  return 0;
}

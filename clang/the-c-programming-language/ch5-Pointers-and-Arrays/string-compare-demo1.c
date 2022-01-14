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

int main(int argc, char const *argv[])
{
  char *s1 = "abc";
  char *t1 = "abd";
  printf("compare result: %d\n", custom_strcmp(s1, t1));

  char *s2 = "abd";
  char *t2 = "abc";
  printf("compare result: %d\n", custom_strcmp(s2, t2));

  char *s3 = "abc";
  char *t3 = "abc";
  printf("compare result: %d\n", custom_strcmp(s3, t3));

  char *s4 = "abc";
  char *t4 = "acc";
  printf("compare result: %d\n", custom_strcmp(s4, t4));

  return 0;
}

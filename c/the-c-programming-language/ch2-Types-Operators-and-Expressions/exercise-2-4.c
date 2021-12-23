#include <stdio.h>

/*
page 39

参考: ch2-Types-Operators-and-Expressions/squeeze_demo.c
重新编写函数 squeeze(char s1[], char s2[]) ，将字符串s1中任何与字符串s2中字符匹配的字符都删除
*/

void squeeze(char s1[], char s2[]){
  int i = 0;
  int j;
  int k = 0;

  for(; s1[i] != '\0'; i++){
    for(j = 0; s2[j] != '\0' && s1[i] != s2[j]; j++){
      ;
    }

    if(s2[j] == '\0'){
      s1[k] = s1[i];
      k++;
    }
  }

  s1[k] = '\0';
}


int main(int argc, char const *argv[])
{
  char str1[10] = "hello";
  char str2[10] = "abc";
  printf("str1: %s\n", str1);
  printf("str2: %s\n", str2);

  printf("\n");
  squeeze(str1, str2);
  printf("str1: %s\n", str1);
  printf("str2: %s\n", str2);

  return 0;
}
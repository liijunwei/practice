#include <stdio.h>

/*
page 39

参考: ch2-Types-Operators-and-Expressions/squeeze_demo.c
重新编写函数 squeeze(char s1[], char s2[]) ，将字符串s1中任何与字符串s2中字符匹配的字符都删除
*/

void squeeze(char s1[], char s2[]){
  int i;
  int j;
  int k;

  for(i = 0, k = 0; s1[i] != '\0'; i++){
    for(j = 0; s2[j] != '\0' && s1[i] != s2[j]; j++){
      ;
    }

    // s2[j] == '\0' 和上面for循环里的 s2[j] != '\0' 条件是矛盾的
    // 即 要么执行for循环, 要么执行这个if条件
    if(s2[j] == '\0'){
      s1[k] = s1[i];
      printf("k: %d s1[k]: %c\n", k, s1[k]);
      printf("i: %d s1[i]: %c\n", i, s1[i]);

      k++;
    }
  }

  s1[k] = '\0';
}


int main(int argc, char const *argv[])
{
  char str1[10] = "hello";
  char str2[10] = "eo";
  printf("str1: %s\n", str1);
  printf("str2: %s\n", str2);

  printf("\n");
  squeeze(str1, str2);
  printf("str1: %s\n", str1);
  printf("str2: %s\n", str2);

  return 0;
}
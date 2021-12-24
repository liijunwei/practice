#include <stdio.h>

/*
page 38

编写函数any(s1,s2), 将字符串s2中的任一字符在字符串s1中第一次出现的位置作为结果返回
如果s1中不包含s2中的字符, 则返回-1(标准库函数strpbrk具有同样的\n功能, 但它返回的是指向该位置的指针)
*/

int any(char s1[], char s2[]){
  int i;
  int j;

  for(i = 0; s1[i] != '\0'; i++){
    for(j = 0; s2[j] != '\0' && s1[i] != s2[j]; j++){
      ;
    }

    if(s2[j] == '\0'){
      // do nothing
    }

    if(s2[j] == s1[i]){
      return i;
    }
  }

  return -1;
}

int main(int argc, char const *argv[])
{
  char str1[10] = "hello";
  char str2[10] = "o";

  printf("%d\n", any(str1, str2));

  return 0;
}
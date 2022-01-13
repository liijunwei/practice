#include <stdio.h>

/*
page 90

*/

// 将指针t指向的字符串复制到指针s指向的位置(使用数组下标实现)
void custom_strcpy(char *s, char *t){
  int i = 0;

  while((s[i] = t[i]) != '\0'){
    i++;
  }
}

int main(int argc, char const *argv[])
{
  char strbuf[100];
  char *strptr = "99 bottles of oop";

  printf("strbuf: %s\n", strbuf);
  printf("strptr: %s\n", strptr);

  custom_strcpy(strbuf, strptr);
  printf("strbuf: %s\n", strbuf);

  return 0;
}

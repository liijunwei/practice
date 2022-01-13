#include <stdio.h>

/*
page 90

*/

// 将指针t指向的字符串复制到指针s指向的位置(使用数组下标实现)
void custom_strcpy(char *s, const char *t){
  while((*s++ = *t++) != '\0'){
    ;
  }
}

int main(int argc, char const *argv[])
{
  char strbuf[100];
  char *strptr = "99 bottles of oop";

  printf("strbuf(before): %s\n", strbuf);

  custom_strcpy(strbuf, strptr);
  printf("strbuf(after): %s\n", strbuf);

  return 0;
}

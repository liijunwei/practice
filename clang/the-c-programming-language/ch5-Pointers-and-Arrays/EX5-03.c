#include <stdio.h>

/*
page 92

用指针的方式实现第2章中函数的 strcat
ch2-Types-Operators-and-Expressions/strcat_demo.c

*/

// 将t指向的字符串复制到s指向的字符串的尾部
void custom_strcat(char s[], char t[]){
  while(*s != '\0'){
    s++;
  }

  while((*s = *t) != '\0'){
    s++;
    t++;
  }
}

int main(int argc, char const *argv[])
{
  char str[100] = "Hello";
  custom_strcat(str, " Xiaoli.");
  custom_strcat(str, " Good");
  custom_strcat(str, " Morning.");
  custom_strcat(str, " 20211223 0813 ");
  custom_strcat(str, " 20220115 1341");

  printf("%s\n", str);

  return 0;
}

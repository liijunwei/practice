#include <stdio.h>

// page 38

// 将字符串t接到字符串s尾部, s必须有足够大的空间
void custom_strcat(char s[], char t[]){
  int i = 0;
  int j = 0;

  while(s[i] != '\0'){
    i++;
  }

  while((s[i++] = t[j++]) != '\0'){
    ;
  }

}

int main(int argc, char const *argv[])
{
  char str[100] = "Hello";
  custom_strcat(str, " Xiaoli.");
  custom_strcat(str, " Good");
  custom_strcat(str, " Morning.");
  custom_strcat(str, " 20211223 0813");

  printf("%s\n", str);

  return 0;
}

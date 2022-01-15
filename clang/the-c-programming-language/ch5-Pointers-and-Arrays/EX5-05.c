#include <stdio.h>
#include <string.h>
#include <assert.h>

/*
page 92

实现库函数 strncpy/strncat/strncmp, 它们最多对参数字符串中的前n个字符进行操作
例如: 函数strncpy(s, t, n)将t中最前n个字符复制到s中
更详细的说明参见附录B


*/

// 将t中最前n个字符复制到s中
// copy n characters from t to s
void custom_strncpy(char *s, char *t, int n){
  while(*t != '\0' && n-- > 0){
    *s++ = *t++;
  }

  while(n-- > 0){
    *s++ = '\0';
  }
}

// 将t中最前n个字符接到到s后面
// concatenate n characters of t to the end of s
void custom_strncat(char *s, char *t, int n){
  custom_strncpy(s + strlen(s), t, n);
}

// 将t中最前n个字符和s进行比较
// compare at most n characters of t with s
int custom_strncmp(char *s, char *t, int n){
  for(; *s == *t; s++, t++){
    if(*s == '\0' || --n <= 0){
      return 0;
    }
  }

  return *s - *t;
}

int main(int argc, char const *argv[])
{
  char str[100];

  custom_strncpy(str, "Roses are red", 5);
  assert(strcmp(str, "Roses") == 0);

  custom_strncpy(str, "Roses are red", 6);
  assert(strcmp(str, "Roses ") == 0);

  custom_strncpy(str, "Roses are red", 13);
  assert(strcmp(str, "Roses are red") == 0);

  return 0;
}

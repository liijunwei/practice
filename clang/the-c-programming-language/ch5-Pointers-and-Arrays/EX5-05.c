#include <stdio.h>
#include <string.h>
#include <assert.h>

/*
page 92

实现库函数 strncpy/strncat/strncmp, 它们最多对参数字符串中的前n个字符进行操作
例如: 函数strncpy(s, t, n)将t中最前n个字符复制到s中
更详细的说明参见附录B

TODO 反复写几次, 看明白

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
  char sample1[100];

  custom_strncpy(sample1, "Roses are red", 0);
  assert(strcmp(sample1, "") == 0);

  custom_strncpy(sample1, "Roses are red", 1);
  assert(strcmp(sample1, "R") == 0);

  custom_strncpy(sample1, "Roses are red", 5);
  assert(strcmp(sample1, "Roses") == 0);

  custom_strncpy(sample1, "Roses are red", 6);
  assert(strcmp(sample1, "Roses ") == 0);

  custom_strncpy(sample1, "Roses are red", 13);
  assert(strcmp(sample1, "Roses are red") == 0);

  char sample2[100] = "Roses are red";
  custom_strncat(sample2, " Mud is fun", 0);
  assert(strcmp(sample2, "Roses are red") == 0);

  char sample3[100] = "Roses are red";
  custom_strncat(sample3, " Mud is fun", 4);
  assert(strcmp(sample3, "Roses are red Mud") == 0);

  char sample4[100] = "Roses are red";
  custom_strncat(sample4, " Mud is fun", 11);
  assert(strcmp(sample4, "Roses are red Mud is fun") == 0);

  assert(custom_strncmp("ABCDEFG", "ABC", 0) == 0);
  assert(custom_strncmp("ABCDEFG", "ABC", 1) == 0);
  assert(custom_strncmp("ABCDEFG", "ABC", 2) == 0);
  assert(custom_strncmp("ABCDEFG", "ABC", 3) == 0);
  assert(custom_strncmp("ABCDEFG", "BC", 2) == -1);
  assert(custom_strncmp("ABCDEFG", "ABB", 2) == 0);
  assert(custom_strncmp("ABCDEFG", "ABB", 3) == 1);

  return 0;
}

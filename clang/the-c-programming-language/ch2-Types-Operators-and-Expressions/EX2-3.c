#include <stdio.h>
#include <ctype.h>
#include <assert.h>

/*
page 37

编写函数htoi(s)，把由十六进制数字组成的字符串(包含可选的前缀0x或0X)转换为与之等价的整型值。字符串中允许包含的数字包括:0~9、a~f以及A~F

OK 参考: ch2-Types-Operators-and-Expressions/atoi.c
*/

int custom_htoi(const char s[]){
  int flag = 1;
  int n = 0;
  int hexdigit;
  int i = 0;

  if(s[i] == '0'){
    ++i;

    if(tolower(s[i]) == 'x'){
      ++i;
    }
  }

  for(; flag == 1; i++){
    if(s[i] >= '0' && s[i] <= '9'){
      hexdigit = s[i] - '0';
    } else if(tolower(s[i]) >= 'a' && tolower(s[i]) <= 'f'){
      hexdigit = tolower(s[i]) - 'a' + 10;
    } else {
      flag = -1;
    }

    if(flag == 1){
      n = 16 * n + hexdigit;
    }
  }

  return n;
}

int main(int argc, char const *argv[])
{
  printf("assertion started...\n");

  assert(custom_htoi("0x1a") == 26);
  assert(custom_htoi("0X1A") == 26);
  assert(custom_htoi("0XFA129") == 1024297);
  assert(custom_htoi("FA129") == 1024297);

  printf("assertion pass!\n");
  return 0;
}


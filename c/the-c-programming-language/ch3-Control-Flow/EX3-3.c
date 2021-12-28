#include <stdio.h>

/*
page 52

编写函数 expand(s1,s2)，将字符串s1中类似于a-z一类的速记符号在字符串s2中扩展为等价的完整列表abc...xyz。
该函数可以处理大小写字母和数字，并可以处理a-b-c、a-z0-9与a-z等类似的情况。作为前导和尾随的字符原样复制
*/

void expand(const char s1[], char s2[]){
  char c;
  int i = 0;
  int j = 0;

  while((c = s1[i]) != '\0'){
    i++;
    if(s1[i] == '-' && s1[i+1] >= c){
      i++;
      while(c < s1[i]){ // expand
        s2[j++] = c++;
      }
    } else {
      s2[j++] = c;
    }
  }

  s2[j] = '\0';
}

int main(int argc, char const *argv[])
{
  char s1[] = "1-9a-fX-Z";
  printf("s1 -> %s\n", s1);
  char s2[100];

  expand(s1, s2);
  printf("s2 -> %s\n", s2);

  return 0;
}

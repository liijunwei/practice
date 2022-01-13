#include <stdio.h>
#include <assert.h>

/*
page 88

求字符串的长度

指针的减法也是有意义的: 如果p和q执行相同数组中的元素, 且p<q, 那么q-p+1就是位于p和q指向的元素之间的元素的数目
*/

int custom_strlen(char *s){
  char *p = s;

  while(*p != '\0'){
    p++;
  }

  return p - s;
}

int main(int argc, char const *argv[])
{

  assert(0 == custom_strlen(""));
  assert(1 == custom_strlen("a"));
  assert(2 == custom_strlen("ab"));
  assert(3 == custom_strlen("abc"));
  assert(4 == custom_strlen("abcd"));

  return 0;
}


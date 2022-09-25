#include <stdio.h>

/*
page 49

编写一个函数
escape(s,t)，将字符串t复制到字符串s中，并在复制过程中将换行符、制表符等不可显示字符分别转换为\n、\t等相应的可显示的转义字符序列。
要求使用switch语句。
再编写一个具有相反功能的函数，在复制过程中将转义字符序列转换为实际字符。

*/

// t 复制到 s
void escape(char s[], char t[]) {
  int i;
  int j;

  for (i = 0, j = 0; t[i] != '\0'; i++) {
    switch (t[i]) {
      putchar(t[i]);
    case '\n':
      s[j++] = '\\';
      s[j++] = 'n';
      break;
    case '\t':
      s[j++] = '\\';
      s[j++] = 't';
      break;
    default:
      s[j++] = t[i];
      break;
    }
  }

  s[j] = '\0';
}

// t 复制到 s
void unescape(char s[], char t[]) {
  int i;
  int j;

  for (i = 0, j = 0; t[i] != '\0'; i++) {
    if (t[i] != '\\') {
      s[j++] = t[i];
    } else {
      switch (t[++i]) {
      case 'n':
        s[j++] = '\n';
        break;
      case 't':
        s[j++] = '\t';
        break;
      default:
        s[j++] = '\\';
        s[j++] = t[i];
        break;
      }
    }
  }

  s[j] = '\0';
}

int main(int argc, char const *argv[]) {
  char str_from[10] = {'h', 'e', 'l', '\t', '\n', '1', 'o', '!', '\0'};
  printf("str_from -> %s\n", str_from);

  char str_to[100];
  escape(str_to, str_from);
  printf("str_to -> %s\n", str_to);

  char unescaped_str[100];
  unescape(unescaped_str, str_to);
  printf("unescaped_str -> %s\n", unescaped_str);

  return 0;
}

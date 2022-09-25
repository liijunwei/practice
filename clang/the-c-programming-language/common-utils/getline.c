// 将行保存到s中, 并返回该行的行数
int custom_getline(char s[], int max) {
  int c;
  int i = 0;

  while (--max > 0 && (c = getchar()) != EOF && c != '\n') {
    s[i] = c;
    i++;
  }

  if (c == '\n') {
    s[i] = c;
    ++i;
  }

  s[i] = '\0';

  return i;
}

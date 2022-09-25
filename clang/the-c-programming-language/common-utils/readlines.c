#define MAXLEN 1000
#define ALLOCSIZE 10000
static char allocbuf[ALLOCSIZE];
static char *allocp = allocbuf;

// 返回指向n个字符的指针(ch5-Pointers-and-Arrays/allocbuf-afree-demo.c)
char *alloc(int n) {
  if (allocbuf + ALLOCSIZE - allocp >= n) {
    allocp += n;
    return allocp - n;
  } else {
    printf("memory not enough(%d > %d)\n", n, ALLOCSIZE);
    return NULL;
  }
}

int readlines(char *lineptr[], int maxlines) {
  int len;
  int nlines;
  char *p;
  char line[MAXLEN];

  nlines = 0;

  while ((len = custom_getline(line, MAXLEN)) > 0) {
    if (nlines >= maxlines || (p = alloc(len)) == NULL) {
      return -1;
    } else {
      line[len - 1] = '\0'; /* replace '\n' with '\0' */
      strcpy(p, line);      /* char* strcpy( char* dest, const char* src ); */
      lineptr[nlines++] = p;
    }
  }

  return nlines;
}

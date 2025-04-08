#include <stdio.h>

// https://www.runoob.com/cprogramming/c-structures.html
struct Books {
  char title[50];
  char author[50];
  char subject[100];
  int book_id;
} book = {"C 语言", "RUNOOB", "编程语言", 123456};

int main() {
  printf("title:   %s\n", book.title);
  printf("author:  %s\n", book.author);
  printf("subject: %s\n", book.subject);
  printf("book_id: %d\n", book.book_id);
}

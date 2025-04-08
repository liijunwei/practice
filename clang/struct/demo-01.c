#include <stdio.h>

// https://www.runoob.com/cprogramming/c-structures.html
typedef struct {
  char title[50];
  char author[50];
  char subject[100];
  int book_id;
} Book;

int main() {
  Book book = {"C 语言", "RUNOOB", "编程语言", 123456};

  printf("title:   %s\n", book.title);
  printf("author:  %s\n", book.author);
  printf("subject: %s\n", book.subject);
  printf("book_id: %d\n", book.book_id);
}

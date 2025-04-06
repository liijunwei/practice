#include <stdio.h>

// https://www.runoob.com/cprogramming/c-structures.html
struct Book {
  char title[50];
  char author[50];
  char subject[100];
  int book_id;
};

int main() {
  struct Book book = {
    title: "C 语言",
    author: "RUNOOB",
    subject: "编程语言",
    book_id: 123456
  };

  printf("title:   %s\n", book.title);
  printf("author:  %s\n", book.author);
  printf("subject: %s\n", book.subject);
  printf("book_id: %d\n", book.book_id);
}

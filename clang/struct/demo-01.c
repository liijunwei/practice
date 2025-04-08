#include <stdio.h>

// https://www.runoob.com/cprogramming/c-structures.html
typedef struct {
  char title[50];
  char author[50];
  char subject[100];
  int book_id;
} Book;

void printf_book1(Book book);
void printf_book2(Book *book);

int main() {
  Book book = {"C 语言", "RUNOOB", "编程语言", 123456};

  printf_book1(book);
  printf_book2(&book);
}

void printf_book1(Book book) {
  printf("printf_book1 sizeof(book): %zu\n", sizeof(book));
  printf("sizeof(book): %zu\n", sizeof(book));
  printf("title:   %s\n", book.title);
  printf("author:  %s\n", book.author);
  printf("subject: %s\n", book.subject);
  printf("book_id: %d\n", book.book_id);
  printf("\n");
}

void printf_book2(Book *book) {
  printf("printf_book2 sizeof(book): %zu\n", sizeof(book));
  printf("title:   %s\n", book->title);
  printf("author:  %s\n", book->author);
  printf("subject: %s\n", book->subject);
  printf("book_id: %d\n", book->book_id);
  printf("\n");
}

#include <stdio.h>

/*

page 71

寄存器(register)声明告诉编译器, 它所声明的变量在程序中使用频率较高
其思想是: 将register变量放在机器的寄存器中, 这样可以使程序更小, 执行速度更快
但编译器可以忽略此选项(???)

*/
int sum(register unsigned m, register unsigned n) { return m + n; }

int main(int argc, char const *argv[]) {

  printf("sum is: %d\n", sum(1, 8));
  return 0;
}

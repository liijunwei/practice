/*
page 136

以实现函数printf的简单版本为例, 介绍如何以可移植的方式编写可处理变长参数表的函数
因为重点在于参数的处理, 所以函数minprintf只处理格式字符串和参数, 格式转换通过调用printf实现

函数printf的正确声明形式为 `int printf(char *fmt, ...)`
其中省略号表示参数表中的参数的数量和类型是可变的, 省略号只能放在参数表的末尾

因为minprintf函数不需要向printf函数那样返回实际输出的字符个数, 因此声明为 `void minprintf(char *fmt, ...)`
编写minprintf函数的关键在于如何处理甚至连名字都没有的参数表
标准头文件<stdarg.h>中定义了一组宏定义, 他们对如何遍历参数表进行了定义; 该头文件的实现因不同的机器而不同, 但提供的接口是一致的

va_list类型用于声明一个变量, 该变量将依次引用各参数. 在函数minprintf中, 我们将该变量成为ap, 意思是"参数指针".
宏va_start将ap初始化为指向第一个无名参数的指针. 在使用ap之前, 该宏必须被调用一次.
参数表必须至少包括一个有名参数, va_start, va_start将最后一个有名参数作为起点

每次调用va_arg, 该函数豆浆返回一个参数, 并将ap指向下一个参数. va_arg使用一个类型名来决定返回的对象类型/指针移动的步长
最后, 必须在函数返回之前调用va_end, 以完成一些必要的清理工作


*/

#include <stdio.h>
#include <stdarg.h>

void minprintf(char *fmt, ...);

int main(int argc, char const *argv[]) {
  minprintf("Roses are red, Violets are blue, Sugar is sweet, And so are you.\n");
  minprintf("two person laughing | %d -> %s -> %f\n", 466666, "nihao", 3.1444444);
  minprintf("constant            | %f\n", 3.14159265354);
  minprintf("say hello to        | %s\n", "Jef");
  minprintf("random string       | %abc\n", "laksdjflkj");

  return 0;
}

/* minimal printf with variable argument list */
void minprintf(char *fmt, ...) {
  va_list ap; /* points to each unnamed argument */
  char *p;
  char *sval;
  int ival;
  double dval;

  va_start(ap, fmt); /* make ap point to first unnamed argument */

  for (p = fmt; *p; p++) {
    if (*p != '%') {
      putchar(*p);
      continue;
    }

    switch (*++p) {
      case 'd':
        ival = va_arg(ap, int);
        printf("%d", ival);
        break;
      case 'f':
        dval = va_arg(ap, double);
        printf("%f", dval);
        break;
      case 's':
        for (sval = va_arg(ap, char *); *sval; sval++) {
          putchar(*sval);
        }
        break;
      default:
        putchar(*p);
        break;
    }

  }

  va_end(ap); /* clean up */
}

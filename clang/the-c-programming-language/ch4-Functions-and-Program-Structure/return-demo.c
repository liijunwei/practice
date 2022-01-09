#include <stdio.h>

/*
page 59

return语句后不一定需要表达式

如果某个函数从一个地方返回时有返回值, 而从另一个地方返回时没有返回值, 该函数并不非法
但这可能是一种出问题的征兆

*/

// warning: non-void function does not return a value in all control paths
int foo(int i){
  if(i > 0){
    return 1;
  }
}

int main(int argc, char const *argv[])
{
  foo(998);

  return 0;
}


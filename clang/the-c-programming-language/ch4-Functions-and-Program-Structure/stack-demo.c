#include <assert.h>
#include <stdio.h>

/*
page 65

stack data structure
*/

#define MAXVAL 5    // 栈val的最大深度
int sp = 0;         // 下一个空闲栈的位置
double val[MAXVAL]; // 值栈

// 把f压入栈中
void push(double f) {
  if (sp < MAXVAL) {
    val[sp] = f;
    int tmp_sp = sp;
    sp++;
    printf("将 %f 压入栈位置: %d, 栈中下一空闲位置: %d\n", f, tmp_sp, sp);
  } else {
    printf("Error: stack(size: %d) full, can't push %g\n", MAXVAL, f);
  }
}

// 弹出并返回栈顶的值
double pop(void) {
  if (sp > 0) {
    int tmp_sp = --sp;
    printf("从栈中弹出元素: %f\n", val[tmp_sp]);
    val[sp] = 0.0;
    return val[sp];
  } else {
    printf("Error: stack empty\n");
    return 0.0;
  }
}

void print_stack() {
  for (int i = MAXVAL - 1; i >= 0; i--) {
    printf("%f\n", val[i]);
  }

  printf("\n");
}

int main(int argc, char const *argv[]) {
  push(1);
  push(2);
  push(3);
  push(4);
  push(5);
  // push(6);

  print_stack();

  pop();
  print_stack();

  pop();
  print_stack();

  // pop();
  // print_stack();

  // pop();
  // print_stack();

  // pop();
  // print_stack();

  return 0;
}

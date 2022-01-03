#include <stdio.h>
#include <assert.h>

/*
page 67

ref: stack-demo.c
在栈操作中添加几个命令, 分别用于:
    + 在不弹出元素的情况下打印栈顶元素
    + 复制栈顶元素
    + 交换栈顶两个元素的值
    + 清空栈

*/

#define MAXVAL 5    // 栈val的最大深度
int sp = 0;         // 下一个空闲栈的位置
double val[MAXVAL]; // 值栈

// 把f压入栈中
void push(double f){
  if(sp < MAXVAL){
    val[sp] = f;
    int tmp_sp = sp;
    sp++;
    printf("将 %f 压入栈位置: %d, 栈中下一空闲位置: %d\n", f, tmp_sp, sp);
  } else {
    printf("Error: stack(size: %d) full, can't push %g\n", MAXVAL, f);
  }
}

// 弹出并返回栈顶的值
double pop(void){
  if(sp > 0) {
    int tmp_sp = --sp;
    printf("从栈中弹出元素: %f\n", val[tmp_sp]);
    val[sp] = 0.0;
    return val[sp];
  } else {
    printf("Error: stack empty\n");
    return 0.0;
  }
}

void print_stack(){
  for(int i = MAXVAL - 1; i >= 0; i--){
    printf("%f\n", val[i]);
  }

  printf("\n");
}

// 在不弹出元素的情况下打印栈顶元素
void peek(){
  if(sp > 0){
    printf("The peek element of the stack is %g\n", val[sp-1]);
  } else {
    printf("error: stack empty\n");
  }
}

// 复制栈顶元素
void duplicate_top(){
  if(sp > 0){
    val[sp] = val[sp - 1];
    printf("Duplicating peek value of %f\n", val[sp - 1]);
    sp++;
  } else {
    printf("error: stack empty\n");
  }
}

void swap(){
  double temp = val[sp - 1];
  val[sp - 1] = val[sp - 2];
  val[sp - 2] = temp;
}

// 交换栈顶两个元素的值
void swap_top_two(){
  if(sp > 1){
    swap();
  } else {
    printf("error: stack has less than two elements\n");
  }
}

    // + 清空栈


int main(int argc, char const *argv[])
{
  push(1);
  push(2);
  push(3);
  push(4);
  // duplicate_top();
  push(6);
  peek();
  swap_top_two();

  print_stack();

  pop();
  print_stack();

  return 0;
}

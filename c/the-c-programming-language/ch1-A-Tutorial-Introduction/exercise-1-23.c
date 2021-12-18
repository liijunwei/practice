#include <stdio.h>

// page 25

/*
编写一个删除C语言程序中所有的注释语句
要正确处理带引号的字符串与字符常量
在C语言程序中，注释不允许嵌套
*/

void rcomment(int c);
void in_comment(void);
void echo_quote(int c);

int main(){
  rcomment(1);
  in_comment();
  echo_quote(1);

  return 0;
}

void rcomment(int c){

}

void in_comment(void){

}

void echo_quote(int c){

}


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

// remove all comments from a valid c program
int main(){
  int c;
  int d;
  while((c = getchar()) != EOF){
    rcomment(c);
  }

  return 0;
}

void rcomment(int c){

}

void in_comment(void){

}

void echo_quote(int c){

}


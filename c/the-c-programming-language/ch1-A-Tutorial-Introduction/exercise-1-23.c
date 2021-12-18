#include <stdio.h>

// page 25

/*
编写一个删除C语言程序中所有的注释语句
要正确处理带引号的字符串与字符常量
在C语言程序中，注释不允许嵌套
*/

// TODO 有问题: 去不掉以 "// " 开头的注释

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

// 原来去除注释, 使用递归实现的吗?

// read each character, remove the comments
void rcomment(int c){
  int d;

  if(c == '/'){
    if((d = getchar()) == '*'){
      in_comment(); // beginning of comment
    } else if(d == '/'){
      putchar(c); // another slash
      rcomment(d); // recursively ... ?
    } else{
      putchar(c); // not a comment
      putchar(d);
    }
  } else if(c == '\'' || c == '"'){
    echo_quote(c); // quote begins
  } else{
    putchar(c); // not a comment
  }
}

// inside of a valid comment
void in_comment(void){
  int c = getchar(); // prev character
  int d = getchar(); // curr character

  while(c != '*' || d != '/'){ // search for end
    c = d;
    d = getchar();
  }
}

// echo characters within quotes
void echo_quote(int c){
  int d;

  putchar(c);
  while((d = getchar()) != c){ // search for end
    putchar(d);
    if(d == '\\'){
      putchar(getchar()); // ignore escape seq
    }
  }

  putchar(d);
}


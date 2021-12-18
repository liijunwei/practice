#include <stdio.h>

// page 25

/*
编写一个程序，查找C语言程序中的基本语法错误，如圆括号、方括号以及花括号不配对等。
要正确处理引号(包括单引号、双引号)、转义字符序列与注释。
*/

/*
测试用的 输入

--- Unbalanced braces
int main(int argc, char const *argv[]){return 0;

--- Unbalanced brackets
int main(int argc, char const *argv[){return 0;}

--- Unbalanced parentheses
int main(int argc, char const *argv[]{return 0;}

*/

int brace; // {}
int brack; // []
int paren; // ()

void in_quote(int c);
void in_comment(void);
void search(int c);

// rudimentary(初级的) syntax checker for C programs
int main(){
  int c;
  extern int brace;
  extern int brack;
  extern int paren;

  while((c = getchar()) != EOF){
    if(c == '/'){
      if((c = getchar()) == '*'){
        in_comment();
      } else{
        search(c);
      }
    } else if(c == '\'' || c == '"'){
      in_quote(c);
    } else{
      search(c);
    }

    if(brace < 0){
      printf("Unbalanced braces\n");
      brace = 0;
    } else if(brack < 0){
      printf("Unbalanced brackets\n");
      brack = 0;
    } else if(paren < 0){
      printf("Unbalanced parentheses\n");
      paren = 0;
    }
  }

  if(brace > 0){
    printf("Unbalanced braces\n");
  } else if(brack > 0){
    printf("Unbalanced brackets\n");
  } else if(paren > 0){
    printf("Unbalanced parentheses\n");
  }

  return 0;
}

// inside quote
void in_quote(int c){
  int d;
  while((d = getchar()) != c){ // search end quote
    if(d == '\\'){
      getchar(); // ignore escape seq
    }
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

void search(int c){
  extern int brace;
  extern int brack;
  extern int paren;

  if(c == '{'){
    ++brace;
  } else if(c == '}'){
    --brace;
  } else if(c == '['){
    ++brack;
  } else if(c == ']'){
    --brack;
  } else if(c == '('){
    ++paren;
  } else if(c == ')'){
    --paren;
  }
}

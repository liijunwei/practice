#include <stdio.h>
#include <ctype.h>
#include <stdlib.h>

/*
page 68

另一种方法是 通过getline函数读入整个输入行, 这种情况下可以不使用getch函数和ungetch函数
请运用这一方法修改计算器程序

*/

// 将行保存到s中, 并返回该行的行数
int custom_getline(char s[], int max){
  int c;
  int i = 0;

  while(--max > 0 && (c = getchar()) != EOF && c != '\n'){
    s[i] = c;
    i++;
  }

  if(c == '\n'){
    s[i] = c;
    ++i;
  }

  s[i] = '\0';

  return i;
}

#define MAXVAL 100 // 栈val的最大深度
int sp = 0; // 下一个空闲栈的位置
double val[MAXVAL]; // 值栈

// 把f压入栈中
void push(double f){
  if(sp < MAXVAL){
    val[sp] = f;
    sp++;
  } else {
    printf("Error: stack full, can't push %g\n", f);
  }
}

// 弹出并返回栈顶的值
double pop(void){
  if(sp > 0) {
    return val[--sp];
  } else {
    printf("Error: stack empty\n");
    return 0.0;
  }
}

#define BUFSIZE 100
char buf[BUFSIZE];
int bufp = 0; // buf中下一空闲的位置

// 取一个字符(可能是压回的字符)
int getch(void){
  return (bufp > 0) ? buf[--bufp] : getchar();
}

// 把字符压回输入中
void ungetch(int c){
  if(bufp >= BUFSIZE) {
    printf("Ungetch, too many characters\n");
  } else {
    buf[bufp++] = c;
  }
}

#define NUMBER '0'
#define MAXLINE 100

int li = 0; // input line index
char line[MAXLINE];

int getop(char s[]){
  int c;
  int i;

  if(line[li] == '\0'){
    if(custom_getline(line, MAXLINE) == 0){
      return EOF;
    } else {
      li = 0;
    }
  }

  while((s[0] = c = line[li++]) == ' ' || c == '\t') {
    ;
  }

  s[1] = '\0';
  if(!isdigit(c) && c != '.'){
    return c; // not a number
  }

  i = 0;
  if(isdigit(c)){ // integer part
    while(isdigit(s[++i] = c = line[li++])){
      ;
    }
  }

  if(c == '.'){ // fraction part
    while(isdigit(s[++i] = c = line[li++])){
      ;
    }
  }

  s[i] = '\0';
  li--;

  return NUMBER;
}

#define MAXOP 100

int main(int argc, char const *argv[])
{
  int type;
  double op2;
  char s[MAXOP];

  while((type = getop(s)) != EOF){
    switch(type){
      case NUMBER:
        push(atof(s));
        break;
      case '+':
        push(pop() + pop());
        break;
      case '*':
        push(pop() * pop());
        break;
      case '-':
        op2 = pop();
        push(pop() - op2);
        break;
      case '/':
        op2 = pop();
        if(op2 != 0.0){
          push(pop() / op2);
        } else {
          printf("Error: zero devisor\n");
        }
        break;
      case '\n':
        printf("Result:\t%.8g\n", pop());
        break;
      default:
        printf("Error: unknown command %s\n", s);
        break;
    }
  }

  return 0;
}

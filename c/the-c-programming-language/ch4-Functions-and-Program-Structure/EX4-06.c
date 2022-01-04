#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <ctype.h>
#include <math.h> /* for fomod */

/*
page 67

给计算机程序增加处理变量的命令(提供26个具有单个英文字母变量名的变量很容易)
增加一个变量存放最近打印的值


*/

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
int bufp = 0;

int getch(void){
  return (bufp > 0) ? buf[--bufp] : getchar();
}

void ungetch(int c){
  if(bufp >= BUFSIZE) {
    printf("Ungetch, too many characters\n");
  } else {
    buf[bufp++] = c;
  }
}

void mathfunc(char s[]){
  double op2;

  if(strcmp(s, "sin") == 0){
    push(sin(pop()));
  } else if(strcmp(s, "cos") == 0){
    push(cos(pop()));
  } else if(strcmp(s, "exp") == 0){
    push(exp(pop()));
  } else if(strcmp(s, "pow") == 0){
    op2 = pop();
    push(pow(pop(), op2));
  } else {
    printf("error: %s is not supported\n", s);
  }
}

#define NUMBER '0' // a number was found
#define NAME   'n' // a name was found

// 获取下一操作符/操作数
int getop(char s[]){
  int i;
  int c;

  while((s[0] = c = getch()) == ' ' || c == '\t'){
    ;
  }

  s[1] = '\0';
  i = 0;

  if(islower(c)){
    while(islower(s[++i] = c = getch())){
      ;
    }

    s[i] = '\0';
    if(c != EOF){
      ungetch(c);
    }

    if(strlen(s) > 1){
      return NAME;
    } else {
      return c;
    }
  }

  if(!isdigit(c) && c != '.') {
    return c; // 不是数
  }

  if(isdigit(c)) {
    // 整数部分
    while(isdigit(s[++i] = c = getch())){
      ;
    }
  }

  if(c == '.') {
    // 小数部分
    while(isdigit(s[++i] = c = getch())){
      ;
    }
  }

  s[i] = '\0';

  if(c != EOF) {
    ungetch(c);
  }

  return NUMBER;
}

#define MAXOP  100

int main(int argc, char const *argv[])
{
  int type;
  int i;
  int var = 0;
  double op2;
  double v;
  double variable[26];
  char s[MAXOP];

  for(i = 0; i < 26; i++){
    variable[i] = 0.0;
  }

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
      case '%':
        op2 = pop();
        if(op2 != 0.0){
          push(fmod(pop(), op2));
        } else {
          printf("Error: zero devisor\n");
        }
        break;
      case NAME:
        mathfunc(s);
        break;
      case '=':
        pop();
        if(var >= 'A' && var <= 'Z'){
          variable[var - 'A'] = pop();
        } else {
          printf("Error: no variable name\n");
        }
        break;
      case '\n':
        printf("Result:\t%.8g\n", pop());
        break;
      default:
        if(type >= 'A' && type <= 'Z'){
          push(variable[type - 'A']);
        } else if(type == 'v') {
          push(v);
        } else {
          printf("Error: unknown command %s\n", s);
        }
        break;
    }

    var = type;
    // printf("Info var is %c\n", var);
  }

  return 0;
}

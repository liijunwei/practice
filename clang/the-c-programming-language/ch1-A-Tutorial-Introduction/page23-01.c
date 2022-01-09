#include <stdio.h>
#include <string.h>

// page 23
#define LENTH_MAX 1000

int max; // 目前为止发现的最长的行
char line[LENTH_MAX]; // 当前的输入行
char longest[LENTH_MAX]; // 最长的行

int custom_getline(void);
void custom_copy(void);

// 注释掉 extern 的行, 是等效的

/*
  某些情况下, extern 可以省略
  在源文件中, 如果外部变量的定义出现在使用它的函数之前, 那么在那个函数中就没有必要使用extern声明
  在通常做法中, 所有的外部变量都声明在文件的开始出, 这样就可以省去extern声明

  人们通常把变量和函数的extern声明放在一个单独的文件中(习惯上称为头文件), 并在每个原文件的开头使用`#include`语句把所有要用的
  头文件包含进来

  后缀名`.h`约定为头文件的拓展名
*/

int main(){
  int len;
  extern int max;
  extern char longest[];

  max = 0;
  while((len = custom_getline()) > 0){
    if(len > max){
      max = len;
      custom_copy();
    }
  }

  if(max > 0){ // 如果存在这样的行
    printf("the longest line is: \n");
    printf("%s", longest);
  }

  return 0;
}

int custom_getline(void){
  int c;
  int i;
  extern char line[];

  for(i = 0; i < LENTH_MAX - 1 && (c = getchar()) != EOF && c != '\n'; ++i){
    line[i] = c;
  }

  if(c == '\n'){
    line[i] = c;
    ++i;
  }

  line[i] = '\0';

  return i;
}

void custom_copy(void){
  int i = 0;
  extern char line[];
  extern char longest[];

  while((longest[i] = line[i]) != '\0'){
    ++i;
  }
}
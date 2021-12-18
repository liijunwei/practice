#include <stdio.h>

// page 25

/*
编写一个程序，把较长的输入行“折”成短一些的两行或多行，折行的位置在输入行的第n列之前的最后一个非空格符之后。
要保证程序能够智能地处理输入行很长以及在指定的列前没有空格或制表符的情况
*/

#define MAXCOL 10 // maximum column of input
#define TABING 8  // tab increment size

char line[MAXCOL];
int exptab(int pos);
int find_blank(int pos);
int new_pos(int pos);
void print_line(int pos);

int main(){
  exptab(1);
  find_blank(1);
  new_pos(1);
  print_line(1);

  return 0;
}

int exptab(int pos){
  return 0;
}

int find_blank(int pos){
  return 0;
}

int new_pos(int pos){
  return 0;
}

void print_line(int pos){

}



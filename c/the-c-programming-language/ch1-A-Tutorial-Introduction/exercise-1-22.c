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

// fold long input lines into two or more shorter lines
int main(){
  int c;
  int pos = 0; // position in the line

  while((c = getchar()) != EOF){
    line[pos] = c;

    if(c == '\t'){
      pos = exptab(pos); // expand tab character
    } else if(c == '\n'){
      print_line(pos); // print current input line
      pos = 0;
    } else if(++pos > MAXCOL){
      pos = find_blank(pos);
      print_line(pos);
      pos = new_pos(pos);
    }
  }

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

// print line until pos column
void print_line(int pos){
  for(int i = 0; i < pos; ++i){
    putchar(line[i]);
  }

  // any chars printed?
  if(pos > 0){
    putchar('\n');
  }
}



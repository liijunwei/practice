#include <stdio.h>

// page 25

/*
编写一个程序，把较长的输入行“折”成短一些的两行或多行，折行的位置在输入行的第n列之前的最后一个非空格符之后。
要保证程序能够智能地处理输入行很长以及在指定的列前没有空格或制表符的情况
*/

// TODO 如何验证程序的正确性

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

// expand tab into blanks
int exptab(int pos){
  line[pos] = ' '; //tab is at least one blank
  for(++pos; (pos < MAXCOL) && (pos % TABING != 0); ++pos){
    line[pos] = ' ';
  }

  if(pos < MAXCOL){ // room left in current line
    return pos;
  } else{           // current line if full
    print_line(pos);
    return 0;       // reest current position
  }
}

// find blank's position
int find_blank(int pos){
  while(pos > 0 && line[pos] != ' '){
    --pos;
  }

  if(pos == 0){     // no blank in the line?
    return MAXCOL;
  } else{           // at least one blank
    return pos + 1; // position after the blank
  }

  return 0;
}

// rearrange line with new position
int new_pos(int pos){
  int i;
  int j;

  if(pos <= 0 || pos >= MAXCOL){ // nothing to rearrange
    return 0;
  } else{
    i = 0;
    for(j = pos; j < MAXCOL; ++j){
      line[i] = line[j];
      ++i;
    }

    return i; // new position in line
  }
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



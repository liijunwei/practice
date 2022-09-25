#include <stdio.h>

// page 25

/*
编写一个程序，把较长的输入行"折"成短一些的两行或多行，折行的位置在输入行的第n列之前的最后一个非空格符之后。
要保证程序能够智能地处理输入行很长以及在指定的列前没有空格或制表符的情况
*/

// TODO 如何验证程序的正确性

#define MAXCOL 10 // maximum column of input
#define TABING 8  // tab increment size

/*
test use sentence:
--------------------
today is a wonderful day isn't it?
TODO 输出很怪, 没看明白...
*/

/*
MAXCOL 指定了输入行的折行位置, 即输入行的第n列
pos 指定了程序在问本行中的当前位置
程序将在输入行的每一处第n列之前对该输入行拆行

这个程序把制表符扩展为空格; 没遇到一个换行符就把此前的输入文本打印出来;
每当变量pos值达到MAXCOL时, 就对输入行进行"折叠"

find_blank从输入行的pos出开始倒退着寻找一个空格(目的是 保持折行的单词完整);
如果找到了一个空格, 就返回紧跟着空格符后面的那个位置的下标; 如果没找到空格,
就返回MAXCOL

print_line 打印输出从位置0到pos-1之间的字符

newpos调整输入行, 把从位置pos开始的字符复制到下一个输出行的开始,
然后再返回pos的新值

*/

char line[MAXCOL];
int exptab(int pos);
int find_blank(int pos);
int newpos(int pos);
void print_line(int pos);

// fold long input lines into two or more shorter lines
int main() {
  int c;
  int pos = 0; // position in the line

  while ((c = getchar()) != EOF) {
    line[pos] = c;

    if (c == '\t') {
      pos = exptab(pos); // expand tab character
    } else if (c == '\n') {
      print_line(pos); // print current input line
      pos = 0;
    } else if (++pos > MAXCOL) {
      pos = find_blank(pos);
      print_line(pos);
      pos = newpos(pos);
    }
  }

  return 0;
}

// expand tab into blanks
int exptab(int pos) {
  line[pos] = ' '; // tab is at least one blank
  for (++pos; (pos < MAXCOL) && (pos % TABING != 0); ++pos) {
    line[pos] = ' ';
  }

  if (pos < MAXCOL) { // room left in current line
    return pos;
  } else { // current line if full
    print_line(pos);
    return 0; // reest current position
  }
}

// find blank's position
int find_blank(int pos) {
  while (pos > 0 && line[pos] != ' ') {
    --pos;
  }

  if (pos == 0) { // no blank in the line?
    return MAXCOL;
  } else {          // at least one blank
    return pos + 1; // position after the blank
  }
}

// rearrange line with new position
int newpos(int pos) {
  int i;
  int j;

  if (pos <= 0 || pos >= MAXCOL) { // nothing to rearrange
    return 0;
  } else {
    i = 0;
    for (j = pos; j < MAXCOL; ++j) {
      line[i] = line[j];
      ++i;
    }

    return i; // new position in line
  }
}

// print line until pos column
void print_line(int pos) {
  for (int i = 0; i < pos; ++i) {
    putchar(line[i]);
  }

  // any chars printed?
  if (pos > 0) {
    putchar('\n');
  }
}

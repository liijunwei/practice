#include <stdio.h>

#define CONST_IN  1 // 在单词内
#define CONST_OUT 0 // 在单词外

#define CONST_CHAR_TAB       '\t'
#define CONST_CHAR_NEWLINE   '\n'
#define CONST_CHAR_SPACE     ' '

// page 15 练习1-11

// TODO 如何用程序测试?(assert)

// 统计输入的 行数/单词数/字符数

/*
OK.1 如何验证程序的正确性?

PASS 测试用例.1 无输入, 应该输出 num_line: 0 num_word: 0 num_char: 0
PASS 测试用例.2 输入1个字符, 应该输出 num_line: 1 num_word: 1 num_char: 2
PASS 测试用例.3 输入2个字符组成一个单词, 应该输出 num_line: 1 num_word: 1 num_char: 3
PASS 测试用例.4 输入2个字符组成两个单词, 在同一行, 应该输出 num_line: 1 num_word: 2 num_char: 4
PASS 测试用例.5 输入2个字符组成两个单词, 在不同行, 应该输出 num_line: 2 num_word: 2 num_char: 4
*/

/*
OK.2 如果程序存在错误, 那么什么类型的输入最可能发现这类错误呢?

那些满足边界条件的输入情况最有助于发现程序的错误

边界条件包括:
+ 没有输入
+ 没有单词, 只要换行符
+ 没有单词, 只有 ' ' '\t' '\n'
+ 每个单词各占一行, 没有 ' ' '\t'
+ 单词出现于文本行行首的情况
+ 单词出现于一串空格之后的情况
*/

int main(){
  int num_line = 0;
  int num_word = 0;
  int num_char = 0;

  int c;
  int state = CONST_OUT;

  printf("please enter chars, `ctrl+d` to exit\n");

  while((c = getchar()) != EOF){
    ++num_char;

    if(c == CONST_CHAR_NEWLINE){
      ++num_line;
    }

    if(c == CONST_CHAR_SPACE || c == CONST_CHAR_NEWLINE || c == CONST_CHAR_TAB){
      state = CONST_OUT;
    }
    else if(state == CONST_OUT){
      state = CONST_IN;
      ++num_word;
    }
    else{
      // do nothing
    }
  }

  printf("num_line: %d num_word: %d num_char: %d \n", num_line, num_word, num_char);
}

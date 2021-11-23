#include <iostream>

using namespace std;
// https://www.bilibili.com/video/BV1bs41197KN?p=99&t=440.7

/*
问题描述:
    在一个有100人的大班级中, 存在多个生日相同的同学的概率很高. 先给出每个学生的学号, 出生月日
    请列出所有生日相同的同学

Input:
    第一行为整数n, 表示n个学生, n < 100;
    此后每行包含一个字符串和两个证书, 分别表示学生的学号(字符串长度小于10)和出生月(1<=m<=12)日(1<=d<=31)
    学号 月 日之间用空格分隔

Output:
    对每组生日相同的学生输出一行;
    其中前两个数字表示月和日, 后面跟着所有在当天出生的学生的学号, 数字和学号之间用空格分隔;
    对所有的输出, 要求按日期从前到后的顺序输出;
    对生日相同的学号, 按输入顺序输出

*/

struct Student{
  char id[10];
  int birthMonth;
  int birthDay;
};

int main(){
  Student students[100] = {
    {"001", 2, 3},
    {"002", 4, 1},
    {"003", 5, 2},
    {"004", 9, 5},
    {"005", 10, 1},
    {"006", 12, 3},
    {"007", 10, 31},
    {"008", 6, 28},
    {"009", 8, 28},
    {"010", 8, 28},
    {"011", 5, 2},
  };
  int i;
  int j;
  int k;
  int flag;
  int count[100] = {0};

  int n = 11;
  cout << n << " students" << endl;
  cout << endl;

  for(int m = 1; m <= 12; m++){
    for(int d = 1; d <= 31; d++){
      flag = 0;
      j = 0;
      for(int i = 0; i < n; i++){
        // TODO 这部分没看懂
        if(students[i].birthMonth == m && students[i].birthDay == d){
          count[++j] = i;
          flag++;
        }
      }

      if(flag > 1){
        cout << m << " " << d << " ";
        for(k = 1; k <= j; k++){
          cout << students[count[k]].id << " ";
        }
        cout << endl;
      }
    }
  }

  cout << endl;
  return 0;
}

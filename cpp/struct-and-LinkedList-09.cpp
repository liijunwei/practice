#include <iostream>

using namespace std;
// https://www.bilibili.com/video/BV1bs41197KN?p=100
/*
链表可以动态得创建

**动态地**申请内存空间
    int *p = new int(1024);
    delete p;

    int *p = new int[4];
    delete []p;

**动态地**建立链表节点
Node *head;
head = new Node;
*/
struct Node{
  int id;
  Node *next;
};

// 链表
int main(){
  Node *head;
  head = new Node;

  Node *temp = head;
  temp->id = 1;
  temp->next = new Node;
  temp = temp->next;


  cout << endl;
  return 0;
}

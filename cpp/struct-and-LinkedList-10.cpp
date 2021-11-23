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

Node *create(){
  Node *head;
  Node *temp;
  int num = 0;
  int n = 0;

  head = new Node;
  temp = head;

  cout << "please enter node id: ";
  cin >> num;

  while(num != -1){
    n++;
    temp->id = num;
    temp->next = new Node;
    temp = temp->next;

    cout << "please enter node id: ";
    cin >> num;
  }

  if(n == 0){
    head = NULL;
  } else {
    temp->next = NULL;
  }

  return head;
}

// 链表
int main(){
  Node *list01 = create();

  cout << endl;
  return 0;
}

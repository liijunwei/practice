#include <iostream>
#include "struct-and-LinkedList-10-def.cpp"

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

int main(){
  Node *list01 = createLinkedListManually();
  cout << endl;
  traverseLinkedList(list01);

  cout << endl;
  return 0;
}


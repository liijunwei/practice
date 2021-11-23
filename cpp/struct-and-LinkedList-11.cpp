#include <iostream>
#include "struct-and-LinkedList-10-def.cpp"

// https://www.bilibili.com/video/BV1bs41197KN?p=101

// 删除链表元素(head)
// 删除链表元素(tail)
// 删除链表元素(middle)

Node *deleteHeadNode(Node *p){
  Node *temp = NULL;
  temp = p;
  p = p->next;
  delete temp;
  return p;
}

Node *deleteTailNode(Node *p){

}

Node *deleteMiddleNode(Node *p){

}

int main(){
  Node *list01 = createLinkedList();
  cout << endl;
  traverseLinkedList(list01);

  list01 = deleteHeadNode(list01);
  traverseLinkedList(list01);

  cout << endl;
  return 0;
}


#include <iostream>
#include "struct-and-LinkedList-10-def.cpp"

// https://www.bilibili.com/video/BV1bs41197KN?p=101

// 删除链表元素(head)
// 删除链表元素(tail)
// 删除链表元素(middle)

Node *deleteMiddleNode(Node *head){
  if(head == NULL){
    return NULL;
  }

  if(head->next == NULL){
    delete head;
    return NULL;
  }

  Node *temp = NULL;
  temp = head;

  



  return NULL;
}

int main(){
  Node *list01 = createLinkedList();
  cout << endl;
  traverseLinkedList(list01);

  list01 = deleteMiddleNode(list01);
  traverseLinkedList(list01);

  cout << endl;
  return 0;
}


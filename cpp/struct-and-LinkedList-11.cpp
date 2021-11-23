#include <iostream>
#include "struct-and-LinkedList-10-def.cpp"

// https://www.bilibili.com/video/BV1bs41197KN?p=101

// 删除链表元素(head)
// 删除链表元素(tail)
// 删除链表元素(middle)

Node *deleteHeadNode(Node *head){
  if(head == NULL){
    return NULL;
  }

  Node *temp = NULL;
  temp = head;
  head = head->next;
  delete temp;
  return head;
}

int main(){
  Node *list01 = createLinkedListManually();
  cout << endl;
  traverseLinkedList(list01);

  list01 = deleteHeadNode(list01);
  traverseLinkedList(list01);

  cout << endl;
  return 0;
}


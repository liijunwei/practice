#include <iostream>
#include "struct-and-LinkedList-10-def.cpp"

// https://www.bilibili.com/video/BV1bs41197KN?p=101

// 插入链表元素(head)
// 插入链表元素(tail)
// 插入链表元素(middle)

Node *insertHeadNode(Node *head, int id){
  Node *temp = NULL;

  temp = new Node;
  temp->id = id;
  temp->next = head;
  head = temp;

  return head;
}

int main(){
  Node *list01 = createLinkedListDemo();
  traverseLinkedList(list01);

  list01 = insertHeadNode(list01, 20);
  traverseLinkedList(list01);

  cout << endl;
  return 0;
}


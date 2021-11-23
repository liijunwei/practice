#include <iostream>
#include "struct-and-LinkedList-10-def.cpp"

// https://www.bilibili.com/video/BV1bs41197KN?p=101

// 插入链表元素(head)
// 插入链表元素(tail)
// 插入链表元素(middle)

Node *insertTailNode(Node *head, int id){
  Node *temp = NULL;
  Node *prev = NULL;
  Node *nodeToInsert = NULL;

  temp = head;
  while(temp != NULL){
    prev = temp;
    temp = temp->next;
  }

  nodeToInsert = new Node;
  nodeToInsert->id = id;
  prev->next = nodeToInsert;
  nodeToInsert->next = NULL;

  return head;
}

int main(){
  Node *list01 = createLinkedListDemo();
  traverseLinkedList(list01);

  list01 = insertTailNode(list01, 20);
  traverseLinkedList(list01);

  cout << endl;
  return 0;
}


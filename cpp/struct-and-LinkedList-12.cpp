#include <iostream>
#include "struct-and-LinkedList-10-def.cpp"

// https://www.bilibili.com/video/BV1bs41197KN?p=101

// 删除链表元素(head)
// 删除链表元素(tail)
// 删除链表元素(middle)

Node *deleteTailNode(Node *head){
  if(head == NULL){
    return NULL;
  }

  if(head->next == NULL){
    head = NULL;
    return head;
  }

  Node *temp = NULL;
  temp = head;

  while(temp->next->next != NULL){
    temp = temp->next;
  }

  Node *lastNode = temp->next;
  temp->next = NULL;
  delete(lastNode);

  return head;
}

Node *deleteMiddleNode(Node *p){
  return NULL;
}

// TODO 这个0怎么处理???
// 1 2 3 4 5 0
// 1 2 3 4 5
int main(){
  Node *list01 = createLinkedList();
  cout << endl;
  traverseLinkedList(list01);

  list01 = deleteTailNode(list01);
  traverseLinkedList(list01);

  cout << endl;
  return 0;
}


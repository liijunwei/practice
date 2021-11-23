#include <iostream>
#include "struct-and-LinkedList-10-def.cpp"

// https://www.bilibili.com/video/BV1bs41197KN?p=101

// 删除链表元素(head)
// 删除链表元素(tail)
// 删除链表元素(middle)

Node *deleteMiddleNode(Node *head, int id){
  if(head == NULL){
    return NULL;
  }

  Node *temp = NULL;
  Node *prev = NULL;
  temp = head;

  while(temp != NULL && temp->id != id){
    prev = temp;
    temp = temp->next;
  }

  if(temp == NULL){
    cout << "Node with id " << id << " not found.\n";
  }
  else{
    prev->next = temp->next;
    delete temp;
  }

  return head;
}

// Segmentation fault is a specific kind of error caused by accessing memory that “does not belong to you.”
int main(){
  Node *list01 = createLinkedListDemo();
  // Node *list01 = createLinkedListManually();
  traverseLinkedList(list01);

  // "66574 segmentation fault  ./a.out"
  list01 = deleteMiddleNode(list01, 20);
  traverseLinkedList(list01);

  cout << endl;
  return 0;
}


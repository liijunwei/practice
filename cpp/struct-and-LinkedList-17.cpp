#include <iostream>
#include "struct-and-LinkedList-10-def.cpp"

// https://www.bilibili.com/video/BV1bs41197KN?p=101

// 插入链表元素(head)
// 插入链表元素(tail)
// 插入链表元素(middle)

Node *insertNodeAfter(Node *head, int id, int data_to_insert){
  Node *temp = NULL;
  Node *prev = NULL;
  Node *nodeToInsert = NULL;

  temp = head;
  while(temp != NULL && temp->id != id){
    prev = temp;
    temp = temp->next;
  }

  if(temp == NULL){
    cout << "Node with id [" << id << "] not found.\n";
  }
  else{
    prev = temp;
    temp = temp->next;

    nodeToInsert = new Node;
    nodeToInsert->id = data_to_insert;
    prev->next = nodeToInsert;
    nodeToInsert->next = temp;
  }

  return head;
}

int main(){
  Node *list01 = createLinkedListDemo();
  traverseLinkedList(list01);

  list01 = insertNodeAfter(list01, 1 , 998);
  list01 = insertNodeAfter(list01, 3 , 998);
  list01 = insertNodeAfter(list01, 5 , 998);
  list01 = insertNodeAfter(list01, 8 , 998);
  traverseLinkedList(list01);

  cout << endl;
  return 0;
}


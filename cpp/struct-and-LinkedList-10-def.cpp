using namespace std;

struct Node{
  int id;
  Node *next;
};

Node *createLinkedListManually(){
  Node *head = NULL;
  Node *temp = NULL;
  int id = 0;
  int n = 0;

  head = new Node;
  temp = head;

  cout << n+1 << " please enter node id: ";
  cin >> id;

  while(id != -1){
    n++;
    temp->id = id;
    temp->next = new Node;
    temp = temp->next;

    cout << n+1 << " please enter node id: ";
    cin >> id;
  }

  if(n == 0){
    head = NULL;
  } else {
    temp->next = NULL;
  }

  return head;
}

void traverseLinkedList(Node *head) {
  if(head == NULL){
    cout << "Null head passed in..." << endl;
    return void();
  }

  cout << "Traverse and print the linked list: " << endl;
  while(head != NULL) {
    cout << head->id << " -> ";
    head = head->next;
  }
  cout << "NULL\n";
}

Node *newNode(int id){
  struct Node *temp = NULL;
  temp = new Node;
  temp->id = id;
  temp->next = NULL;

  return temp;
}

Node *createLinkedListDemo(){
  struct Node *head = NULL;
  head = newNode(1);
  head->next = newNode(2);
  head->next->next = newNode(3);
  head->next->next->next = newNode(4);
  head->next->next->next->next = newNode(5);

  return head;
}

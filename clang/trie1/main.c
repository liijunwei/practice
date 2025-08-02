#include <assert.h>
#include <stdio.h>

#define MAX_CHILDREN_COUNT 256

typedef struct Node Node;
struct Node {
  // char x;
  Node *children[MAX_CHILDREN_COUNT];
};

#define NODE_POOL_CAP 1024
Node node_pool[NODE_POOL_CAP] = {0};
size_t node_pool_count = 0;

Node *alloc_node(void) {
  assert(node_pool_count < NODE_POOL_CAP);
  return &node_pool[node_pool_count++];
}

void insert_text(Node *root, char *text) {
  if (text == NULL || *text == '\0') {
    return;
  }
  assert(root != NULL);
  size_t index = (size_t)*text;
  if (root->children[index] == NULL) {
    root->children[index] = alloc_node();
  }
  insert_text(root->children[index], text + 1);
}

int main(void) {
  Node *root = alloc_node();
  insert_text(root, "hello");
  return 0;
}

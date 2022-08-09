#include <stdio.h>
#include <stdlib.h>

// p348

typedef int data_t;
// typedef long data_t;
// typedef float data_t;
// typedef double data_t;

typedef struct {
  long len;
  data_t *data;
} vec_rec, *vec_ptr;

/* Create vector of specified length */
vec_ptr new_vec(long len) {
  vec_ptr result = (vec_ptr) malloc(sizeof(vec_rec));
  data_t *data = NULL;

  if(!result) {
    return NULL;
  }

  result->len = len;

  if(len > 0) {
    data = (data_t *)calloc(len, sizeof(data_t));

    if(!data) {
      free((void *) result);
      return NULL;
    }
  }

  result->data = data;
  return result;
}

int get_vec_element(vec_ptr v, long index, data_t *dest) {
  if(index < 0 || index >= v->len) {
    return 0;
  }

  *dest = v->data[index];
  return 1;
}

long vec_length(vec_ptr v) {
  return v->len;
}

int main(int argc, char const *argv[]) {


  return 0;
}

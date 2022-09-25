#include <stdio.h>
#include <stdlib.h>

// p348

// typedef int data_t;
typedef long data_t;
// typedef float data_t;
// typedef double data_t;

typedef struct {
  long len;
  data_t *data;
} vec_rec, *vec_ptr;

/* Create vector of specified length */
vec_ptr new_vec(long len) {
  vec_ptr result = (vec_ptr)malloc(sizeof(vec_rec));
  data_t *data = NULL;

  if (!result) {
    return NULL;
  }

  result->len = len;

  if (len > 0) {
    // man calloc
    data = (data_t *)calloc(len, sizeof(data_t));

    if (!data) {
      free((void *)result);
      return NULL;
    }
  }

  result->data = data;
  return result;
}

int get_vec_element(vec_ptr v, long index, data_t *dest) {
  if (index < 0 || index >= v->len) {
    return 0;
  }

  *dest = v->data[index];
  return 1;
}

long vec_length(vec_ptr v) { return v->len; }

// 计算向量元素的和
#define IDENT 0
#define OP +

// 计算向量元素的乘积
// #define IDENT 1
// #define OP    *

void combine1(vec_ptr v, data_t *dest) {
  long i;

  *dest = IDENT;

  for (i = 0; i < vec_length(v); i++) {
    data_t val;
    get_vec_element(v, i, &val);
    *dest = *dest OP val;
  }
}

void combine2(vec_ptr v, data_t *dest) {
  long i;
  long length = vec_length(v);

  *dest = IDENT;

  for (i = 0; i < length; i++) {
    data_t val;
    get_vec_element(v, i, &val);
    *dest = *dest OP val;
  }
}

// p354
data_t *get_vec_start(vec_ptr v) { return v->data; }

// 令人吃惊的是 这个优化没能显著提升程序性能(问题: 怎么测?), 说明 combine2
// 里反复的边界检查不会让性能变差, 之后(5.11.2)会再次回到这个问题
void combine3(vec_ptr v, data_t *dest) {
  long i;
  long length = vec_length(v);
  data_t *data = get_vec_start(v);

  *dest = IDENT;

  for (i = 0; i < length; i++) {
    *dest = *dest OP data[i];
  }
}

int main(int argc, char const *argv[]) {
  vec_ptr vec_demo1 = new_vec(10);
  vec_demo1->data[0] = 0;
  vec_demo1->data[1] = 1;
  vec_demo1->data[2] = 2;
  vec_demo1->data[3] = 3;
  vec_demo1->data[4] = 4;
  vec_demo1->data[5] = 5;
  vec_demo1->data[6] = 6;
  vec_demo1->data[7] = 7;
  vec_demo1->data[8] = 8;
  vec_demo1->data[9] = 9;

  data_t dest1;
  combine1(vec_demo1, &dest1);
  printf("dest1 is %ld\n", dest1);

  data_t dest2;
  combine2(vec_demo1, &dest2);
  printf("dest2 is %ld\n", dest2);

  data_t dest3;
  combine3(vec_demo1, &dest3);
  printf("dest3 is %ld\n", dest3);

  return 0;
}

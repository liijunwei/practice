#include <stdio.h>

typedef struct redisObject {
    int refcount;
} robj;

int main(){
  robj *o = malloc(sizeof(*o));
  o->refcount = 1;

  printf("%lu \n", sizeof(*o));
}

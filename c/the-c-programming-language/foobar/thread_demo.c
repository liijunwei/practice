#include <stdio.h>
#include <pthread.h>

void *entry_point(void *value){
  printf("hello from the 2nd thread :)\n");

  int *num = (int *)value;
  printf("the value of value is %d\n", *num);

  return NULL;
}

int main(int argc, char const *argv[])
{
  pthread_t thread;

  printf("hello from the 1st thread :D\n");

  int num = 9981;
  pthread_create(&thread, NULL, entry_point, &num);
  pthread_join(thread, NULL);

  return 0;
}

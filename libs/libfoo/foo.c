#include <stdio.h>  // printf
#include <stdlib.h> // getenv
#include <unistd.h> // usleep
#include "libfoo.h" // only to check func sig

int InitLibFoo(void)
{
  const char* g = getenv("FOO");
  if (g == NULL) {
    return 111;
  }
  printf("%s\n", g);
  return 0;
}

void DoThings(void){
  printf("going to wait...");
  sleep(3);
  printf("hello world\n");
}

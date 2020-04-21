#include <stdio.h>

#include "../libfoo/libfoo.h"

int main(void){
  int x = InitLibFoo();
  printf("x=%d\n", x);

  printf("going to do things\n");
  DoThings();
  printf("things done\n");
}

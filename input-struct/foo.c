#include "_cgo_export.h"
#include "type.h" // handwritten c struct
int run(Foo *foo){ // need to hand-write the c struct
  int ret = foo->x+foo->y+foo->z;
  foo->x = foo->z;
  foo->y *= 2;
  foo->z = foo->x;
  
  return ret;
}

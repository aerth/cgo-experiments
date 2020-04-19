#include "_cgo_export.h"
#include <assert.h>
#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

int callFoo(void) {
  // auto-generated result struct, named: '{funcname}_return'
  struct foo_return result = foo(1);
  if (strncmp(result.r0.p, "nonzero input", result.r0.n) != 0) {
    return 111;
  }
  if (result.r1) { // expecting false (zero)
    return 222;
  }
  result = foo(0);
  if (strncmp(result.r0.p, "it works", result.r0.n) != 0) {
    return 333;
  }
  if (!result.r1) { // expecting true (nonzero)
    return 444;
  }
  return 0;
}

#include "app.h"
#include <string.h>

char* getstr(void){
  char* ptr = (char*) malloc (1024 *sizeof(char));
  strcpy(ptr, "dont forget to free() me!");
  return ptr;
}

#include "app.h"
#include <stdio.h>

int algorithm_a(char *buf, uint32_t len) {
  printf("char*:\nreceived: ");
  for (int i = 0; i < len; i++) {
    printf("%02x ", (unsigned char)(buf[i]++));
  }
  printf("\n");
  return 0;
}
int algorithm_b(uint8_t *buf, uint32_t len) {
  printf("uint8_t*:\nreceived: ");
  for (int i = 0; i < len; i++) {
    printf("%02x", buf[i]--);
  }
  printf("\n");
  return 0;
}

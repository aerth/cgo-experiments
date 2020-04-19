package main

//#include "app.h"
import "C"

import (
	"fmt"
	"unsafe"
)

func main() {
	var ptr = C.getstr()
	fmt.Printf("%s\n", C.GoString(ptr))
	C.free(unsafe.Pointer(ptr))
}

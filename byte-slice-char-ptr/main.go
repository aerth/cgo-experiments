package main

//#include "app.h"
import "C"
import (
	"fmt"
	"log"
	"unsafe"
)

func print_hex(b []byte) {
	fmt.Printf("% x\n", b)
}

func main() {
	var l = uint32(16)
	// the underlying buffer, "owned" by go runtime
	buf := make([]byte, l)
	for i := range buf {
		buf[i] = 0x55
	}
	print_hex(buf)
	println()

	// note, c doesn't know the length of the go array so we must tell it
	var ptrChar = (*C.char)(unsafe.Pointer(&buf[0]))
	if e := C.algorithm_a(ptrChar, C.uint(l)); e != 0 {
		log.Fatalln("non zero exit code (1)", e)
	}
	C.fflush(C.stdout)

	fmt.Printf("Go got: ")
	print_hex(buf)
	println()

	// note, c doesn't know the length of the go array so we must tell it
	var ptrUnsigned = (*C.uchar)(unsafe.Pointer(&buf[0]))
	if e := C.algorithm_b(ptrUnsigned, C.uint(l)); e != 0 {
		log.Fatalln("non zero exit code (2)", e)
	}
	C.fflush(C.stdout)
	fmt.Printf("Go got: ")
	print_hex(buf)
	println()

}

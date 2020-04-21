package main

/*
#cgo LDFLAGS: -Llibfoo -lfoo
#cgo CFLAGS: -Ilibfoo -Wall -pedantic

#include "libfoo.h"
*/
import "C"
import "fmt"

func main() {
	fmt.Println("initialized libfoo:", C.InitLibFoo())
	done := make(chan struct{})
	go func() {
		C.DoThings()
		fmt.Println("c got back?")
		done <- struct{}{}
	}()
	fmt.Println("waiting for c...")
	<-done
	fmt.Println("all done.")
}

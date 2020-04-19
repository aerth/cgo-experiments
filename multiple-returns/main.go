package main

//int callFoo(void);
import "C"
import (
	"fmt"
	"log"
)

func main() {
	log.SetFlags(0)
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in foo():", r)
		}
	}()
	if x := C.callFoo(); x != 0 {
		log.Fatalln("failed:", x)
	}
	log.Println("return values handled")
}

//export foo
func foo(i int) (string, bool) {
	if i == 0 {
		return "it works", true
	}
	return "nonzero input", false
}

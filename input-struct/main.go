package main

/*
#include "type.h" // C.Foo
*/
import "C"

import "fmt"

// Foo with x,y,z int won't work in CGO (yet?)
type Foo struct {
	c C.Foo
}

func main() {
	ptr := &C.Foo{1, 2, 3}
	fmt.Println("Struct:", *ptr)
	fmt.Println("Output:", C.run(ptr))
	fmt.Println("Struct:", *ptr)
	fmt.Println("-----")
	// same thing, with nonptr (need to send run() the ptr)
	st := Foo{C.Foo{4, 5, 6}}
	fmt.Println("Struct:", st.c)
	fmt.Println("Output:", C.run(&st.c))
	fmt.Println("Struct:", st.c)
}

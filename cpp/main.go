package main

/*
#cgo LDFLAGS: libfoo.a
void say(char* str);
void sayplusplus(char* str);
int get();
int getplusplus();
*/
import "C"
import "fmt"

func main() {
	fmt.Println("hello from go")
	cstr := C.CString("hello from C")
	cppstr := C.CString("hello from C++")
	C.say(cstr)
	C.sayplusplus(cppstr)

	x := C.get()
	y := C.getplusplus()

	fmt.Printf("C:   %d\nC++: %d\n", x, y)
}

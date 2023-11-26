package main

// void myprint(char * s) {
//	printf("%s\n", s);
// }
import "C"
import "unsafe"

func main() {
	cs := C.CString("Hello from Go!")
	C.myprint(cs)
	C.free(unsafe.Pointer(cs))
}

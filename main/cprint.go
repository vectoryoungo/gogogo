package main
/*
#include <stdio.h>
#include <stdlib.h>
*/
import "C"
import "unsafe"

//test go code call c function
func main() {
	cstr := C.CString("Hello world")
	C.puts(cstr)
	C.free(unsafe.Pointer(cstr))
}

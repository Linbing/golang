package main

/*
#include <stdio.h>
#include <stdlib.h>
#if 1
void hello()
{
	printf("hello world \n");
//return;
}
#endif
*/
import "C"
import (
	"fmt"
	"unsafe"
)

func Prints(s string) {
	p := C.CString(s)
	C.fputs(p, (*C.FILE)(C.stdout))
	C.free(unsafe.Pointer(p))
}
func start_goc() {
	fmt.Println("------start goc-------")
	//fmt.Println(int(C.random()))
	//C.hello()
	//C.printf("in c hello world \n");//error : unexpected type: ...
	s := "linbing"
	Prints(s)

}

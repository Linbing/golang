package main

/*
#include <stdlib.h>
*/
import "C"
import (
	"fmt"
)

func start_goc() {
	fmt.Println("-------------")
	C.printf("c hello world \n")
}

package main

import (
	"fmt"
)

func start_defer() {
	fmt.Println(" >> start defer test ")
	call_test()
	fmt.Println(" return from call_test nomally")
}

func call_test() {
	fmt.Println(" >> in call_test << ")
	defer func() {
		//		if r := recover(); r != nil {
		//			fmt.Println(" >> reconver in call_test ", r)
		//		}
		fmt.Println(" defer in call_test ")
	}()
	fmt.Println(" >> start call_panic ....")
	call_panic(0)
	fmt.Println(" >> return nomally from call_panic")
}

func call_panic(i int) {
	fmt.Println(" >> in call_panic ")
	if i > 3 {
		fmt.Println(" Panicking ....!")
		panic(fmt.Sprintf("%v", i))
	}

	defer fmt.Println(" >> defer in call_panic ", i)
	fmt.Println("call_panic i is : ", i)
	call_panic(i + 1)
}

package main

import (
	"fmt"
	"time"
)

func start_goruntime() {
	var step int = 0
	for {
		time.Sleep(1000 * time.Millisecond)
		fmt.Println("in go runtime")
		step++
		if step > 10 {
			break
		}
	}
}

package main

import (
	"fmt"
	//"math/rand"
	"sync"
)

func start_Waitchannel() {
	fmt.Println("---------")
	var waitGroup sync.WaitGroup
	ch1 := make(chan int, 3)
	ch2 := make(chan int, 3)
	ch3 := make(chan int, 3)
	waitGroup.Add(3)
	go func() {
		for n := range ch1 {
			if n%2 == 0 {
				ch2 <- n
			} else {
				fmt.Printf("Filter %d..[Filter 1]\n", n)
			}
		}
		close(ch2)
		waitGroup.Done()
	}()
	go func() {
		for n := range ch2 {
			if n%5 == 0 {
				ch3 <- n
			} else {
				fmt.Printf("Filter %d ..[Filter 2]\n", n)
			}
		}
		close(ch3)
		waitGroup.Done()
	}()
	go func() {
		for n := range ch3 {
			fmt.Println(n)
		}
		waitGroup.Done()
	}()

	for i := 0; i < 100; i++ {
		//ch1 <- rand.Int63n(100)
		ch1 <- i
	}
	close(ch1)
	waitGroup.Wait()
}

package main

import (
	"fmt"
	"time"
)

func Count(ch chan int) {
	ch <- 12
	fmt.Println("counting ")
}

func t1() {
	fmt.Println("-----t1---------")
	//chs := make([]chan int, 10)
	chs := make(chan int, 10)
	for i := 0; i < 10; i++ {
		//chs[i] = make(chan int)
		//fmt.Println("--------")
		//go Count(chs[i])
		go Count(chs)
	}
	time.Sleep(time.Second * 3)
	close(chs)
	for ch := range chs {
		fmt.Println(ch)
	}
	//
	test := make([]chan int, 10)
	test[1] = make(chan int)
	//	test[1] <- 1
	//	<-test[1]
	//close(test[1])
	//	for i, tt := range test {
	//		fmt.Println(i, "---", tt)
	//	}
	fmt.Println(" out ")
}

func Sendmessage(message chan string, num int) {
	message <- fmt.Sprintf("message %d", num)

}

func t2() {
	message := make(chan string)
	count := 3
	for i := 1; i <= count; i++ {
		fmt.Println("send message")
		go Sendmessage(message, i)

	}
	//	go func() {
	//		for i := 1; i <= count; i++ {
	//			fmt.Println("send message")
	//			message <- fmt.Sprintf("message %d", i)
	//		}
	//		close(message)
	//	}()
	time.Sleep(time.Second * 1)
	//close(message)
	for i := 1; i <= count; i++ {
		fmt.Println(<-message)
	}
	//	for msg := range message {
	//		fmt.Println(msg)
	//	}
}
func t3() {
	c := make(chan int, 2)
	c <- 42
	c <- 34
	close(c)
	//	val := <-c
	//	fmt.Println(val)
	for v := range c {
		fmt.Println(v)
	}

}
func t4() {
	c := make(chan int)
	go func() {
		c <- 42
	}()
	val := <-c
	fmt.Println(val)
}
func t5() {
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue)
	for elem := range queue {
		fmt.Println(elem)
	}
}
func start_channel() {
	//t1()
	//t2()
	//t3() //cause deaklock!
	//t4()
	//t5()
}

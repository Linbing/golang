package main

import (
	"fmt"
	"time"
)

func f1(c1 chan int) {
	c1 <- 1
}
func f2(c2 chan int) {
	c2 <- 1
}
func f3(c3 chan int) {
	time.Sleep(time.Second * 5)
	c3 <- 4

}
func s1() {
	fmt.Println("----")
	c1 := make(chan int)
	c2 := make(chan int)
	c3 := make(chan int)
	//c4 := make(chan int)
	go f1(c1)
	go f2(c2)
	go f3(c3)
	for {
		select {
		case tmp := <-c1:
			fmt.Println("c1----", tmp) //must be io operation
		case tmp := <-c2:
			fmt.Println("c2---", tmp)
		case tmp := <-c3:
			fmt.Println("c3----", tmp)
		case c1 <- 1:
			fmt.Println(" yes ")
		default:
			time.Sleep(time.Second * 1)
			fmt.Println("waitinf")
		}
	}
}

func s2() {
	cs := make(chan int, 3)
	go f1(cs)
	go f2(cs)
	go f3(cs)
	for {
		tmp := <-cs
		//time.Sleep(time.Second * 1)
		fmt.Println("tmp is : ", tmp)
		select {
		case cs <- 1:
			fmt.Println("yes 1")
			time.Sleep(time.Second * 1)
		case cs <- 2:
			fmt.Println("yes 2")
		default:
			time.Sleep(time.Second * 1)
			fmt.Println("waiting")
		}
	}
}

func s3() {
	flag := 0
	ch := make(chan int)
	timeout := make(chan bool, 1)
	go func() {
		time.Sleep(time.Second * 5)
		timeout <- true
	}()
	for {
		select {
		case <-ch:
			fmt.Println("read from ch")
		case <-timeout:
			fmt.Println("ch hase be time out")
			flag = 1
			//break
		default:
			time.Sleep(time.Second * 1)
			fmt.Printf("waitine\n")
		}
		if flag == 1 {
			break
		}
	}
}

func start_select() {
	//s1()
	//	s2()
	s3()
}

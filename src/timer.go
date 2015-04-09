package main

import (
	"fmt"
	"time"
)

func Testtime() {
	fmt.Println("--------")
	timer := time.NewTimer(time.Second * 2)
	<-timer.C
	fmt.Println("timer has been expired")
}

func TestTicke() {
	ticker := time.NewTicker(time.Second)
	for t := range ticker.C {
		fmt.Println("Tick at : ", t)
	}
}

func TimeAndTicker() {
	ticker := time.NewTicker(time.Second)
	go func() {
		for t := range ticker.C {
			fmt.Println(t)
		}
	}()
	timer := time.NewTimer(time.Second * 5)
	<-timer.C
	fmt.Println("timer expired!")
}

func start_timer() {
	//Testtime()
	//TestTicke()
	TimeAndTicker()
}

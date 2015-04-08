package main

import (
	"fmt"
	"mymath"
	"obj"
	"time"
)

type Notifier interface {
	Notify() error
}

//func SendNotification(notify *Notifier) error {
//	return notify.Notify()
//}
func main() {
	fmt.Printf(">> start main << \n")
	fmt.Printf("add = %d ...\n", mymath.Add(1, 2))

	obj1 := &obj.Obj{"linbing", 123}
	fmt.Printf("befor age = %d...\n", obj1.Age)
	obj1.ModifyAge(45)
	//	obj1.Notify()
	//	SendNotification(obj1)
	//notify := obj1
	//var notify Notifier = new(obj.Obj)

	//notify.Notify()
	fmt.Printf("-------------------\n")
	fmt.Printf(" >>   go to music test << \n")
	//start_music()
	//start_array()
	//start_map()
	go start_goruntime()
	var step int = 0
	for {
		time.Sleep(1000 * time.Millisecond)
		fmt.Println("in main ")
		step++
		if step > 10 {
			break
		}
	}
}

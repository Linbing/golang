package main

import "fmt"

func modify(arr []int) { //if arr := [10]int{0,1,2...} param must be arr [10]int,
	arr[0] = 12 //then arr[0]may be modify to 12,else arr[0] is no change
}

func start_array() {
	var arr = []int{1, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	//arr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println("len ", len(arr))
	modify(arr)

	var mySlice []int = arr[:5] //0,1..4, all[:], [5:]5,6,...
	for i, v := range mySlice {
		fmt.Println("i:", i, "value:", v)
	}

	//dSlice := make([]int, 5)
	ddSlice := make([]int, 5, 10)
	fmt.Println(len(ddSlice))
	fmt.Println(cap(ddSlice)) //5,10
	ddSlice = append(ddSlice, 1, 2, 3)
	fmt.Println(len(ddSlice))
	fmt.Println(cap(ddSlice))              //8,10
	tmpSlice := make([]int, 2)             //如果是2，则刚好是10，如果重新分配cap，如果》2,则重新分配cap，比如5，则cap为20
	ddSlice = append(ddSlice, tmpSlice...) //... must be needed
	fmt.Println(len(ddSlice))
	fmt.Println(cap(ddSlice)) //10,10
}

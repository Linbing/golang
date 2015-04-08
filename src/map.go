package main

import "fmt"

type PersonInfo struct {
	Id   int
	Name string
}

func start_map() {
	var personDB map[string]PersonInfo
	personDB = make(map[string]PersonInfo)
	personDB["123"] = PersonInfo{123, "linbing"}

	personDB["456"] = PersonInfo{456, "xiaoyue"}

	person, ok := personDB["456"]
	if ok {
		fmt.Println("Found Persion", person.Name)
	} else {
		fmt.Println("did not found persion")
	}
}

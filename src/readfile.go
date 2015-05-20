package main

import (
	"fmt"
	//"io/ioutil"
)

func start_readfile() {
	fmt.Println(" >> start readfile <<")
	configfile := GetInputFile()
	fmt.Println("configfile ", *configfile)
	//	data, err := ioutil.ReadFile(configfile)
}

package main

import (
	"fmt"
	"os/exec"
)

func start_exec() {
	//cmd := exec.Command("ping", "127.0.0.1")
	cmd := exec.Command("ls", "-alh")
	fmt.Println("in heere")
	out, err := cmd.Output()
	fmt.Println(" here ")
	if err != nil {
		fmt.Println("command error ! ", err.Error())
		return
	}

	fmt.Println(string(out))
}

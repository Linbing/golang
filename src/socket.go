package main

import (
	"fmt"
	"net"
)

func echo() {
	fmt.Println("   start echo ")
	listenr, err := net.Listen("tcp", "0.0.0.0:6666")

}

func start_socket() {
	echo()
}

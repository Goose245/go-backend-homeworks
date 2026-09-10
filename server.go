package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	PORT := ":8080"
	listener, err := net.Listen("tcp", PORT)
	if err != nil {
		fmt.Println("Error on listening!")
		os.Exit(1)
	}
	defer listener.Close()

	for {
		con, err := listener.Accept()
		if err != nil {
			fmt.Println("Error on accepting!")
		}
		go Process(con)
	}
}

func Process(con net.Conn) {
	defer con.Close()
	con.Write([]byte("OK\n"))
}

package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	con, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error on connecting")
		os.Exit(1)
	}
	defer con.Close()
	response, err := bufio.NewReader(con).ReadString('\n')
	if err != nil {
		fmt.Println("Error on catching feedback")
		os.Exit(1)
	}
	if response != "OK\n" {
		os.Exit(1)
	}
}

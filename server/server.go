package main

import (
	"fmt"
	"net"
)


func main() {

	ln,err :=  net.Listen("tcp",":8080")

	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		conn,err := ln.Accept()

		if err != nil {
			fmt.Println(err)
			return
		}
		// Handle the connection in a new goroutine
        go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {

	defer conn.Close()
	buf := make([]byte,1024)

	_,err := conn.Read(buf)
	if err != nil {
		fmt.Println(err)
			return
	}
	 // Print the incoming data
	 fmt.Printf("Received: %s", buf)
}
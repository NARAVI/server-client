package main

import (
	"fmt"
	"net"
)

func main() {

	conn,err := net.Dial("tcp",":8080")
	if err != nil {
        fmt.Println(err)
        return
    }

	_,err = conn.Write([]byte("Hello server"))
	if err != nil {
        fmt.Println(err)
        return
    }
	conn.Close()
}
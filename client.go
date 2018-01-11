package main

import (
	"fmt"
	"net"
	// "os"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8000")
	// defer conn.Close()

	if(err != nil) {
		fmt.Println(err)
	}
	var r string
	for {
		
		fmt.Scan(&r)
		conn.Write([]byte(r))
		
		fmt.Println(r)
		if(r == "exit") {
		// 	conn.Close()
		 	return
		}
	}
}
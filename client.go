package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8000")
	defer conn.Close()

	if err != nil {
		fmt.Println(err)
	}
	var r string
	for {

		fmt.Scan(&r)

		fmt.Println(r)
		switch string(r) {
		case "exit":
			return
		default:
			conn.Write([]byte(r))
		}
	}
}

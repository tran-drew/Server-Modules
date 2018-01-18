package main

import (
	"fmt"
	"io"
	"net"
)

// "os")

func main() {
	fmt.Println("Hello World")

	//Create a Listener
	listener, err := net.Listen("tcp", "127.0.0.1:8000")
	defer listener.Close()

	//If error returned when creating Listener
	if err != nil {
		fmt.Println(err)
	}

	//Wait for connection; store into b
	b, err := listener.Accept()
	defer b.Close()

	for {
		str := make([]byte, 256)
		n, err := b.Read(str)
		data := str[:n]

		go func(string) {
			fmt.Println(string(data))
		}(string(data))
		if err == io.EOF {
			return
		}
	}

}

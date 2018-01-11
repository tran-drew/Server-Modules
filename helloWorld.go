package main
import ("fmt"
		"net")

func main() {
	fmt.Println("Hello World")

	listener, err := net.Listen("tcp", "127.0.0.1:8000")

	fmt.Print(listener.Addr())

	if(err != nil) {
		fmt.Println(err)
	}

	listener.Accept()

	fmt.Println("Success!")
}
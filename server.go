package main
import ("fmt"
		"net")
		// "os")

func main() {
	fmt.Println("Hello World")

	//Create a Listener
	listener, err := net.Listen("tcp", "127.0.0.1:8000")
	// defer listener.Close()

	//If error returned when creating Listener
	if(err != nil) {
		fmt.Println(err)
	}
	
	//Wait for connection; store into b
	b, err := listener.Accept()
	// defer b.Close()

	for {
		str := make([]byte, 256)
		b.Read(str)
		fmt.Println(string(str))

		if(string(str) == "\n") {
			fmt.Println("Took an exit")
		}
	}

}
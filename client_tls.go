package main


import (
	"fmt"
	"log"
	"net"
	tls "github.com/Jigsaw-Code/utls"
)

func main() {
	dialConn, err := net.Dial("tcp", "localhost:16667")
	if err != nil {
		log.Fatalln(err)
	}

	config := tls.Config{ServerName: "www.google.com"}
	tlsConn := tls.Client(dialConn, &config)
	defer tlsConn.Close()

	// _, err = dialConn.Write([]byte("Hello, Server!"))
	_, err = tlsConn.Write([]byte("Hello, Server!"))
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Message sent: Hello Server!")

	// _, err = dialConn.Write([]byte("How are you?"))
	_, err = tlsConn.Write([]byte("Hello, Server!"))
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Message sent: How are you?")

	for {
		buffer := make([]byte, 14096)
		dataSize, err := dialConn.Read(buffer)
		tlsConn.Read(buffer)
		if err != nil {
			fmt.Println("connection closed")
			return
		}

		data := buffer[:dataSize]
		fmt.Println("received message: ", string(data))
	}
}
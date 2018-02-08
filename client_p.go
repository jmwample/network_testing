package main
// Parallel
// 		TCP packets are combined by the echo server because no application
// 		is there to separate them. Same on the return journey.


import (
	"fmt"
	"log"
	"net"
	"strconv"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:16667")
	if err != nil {
		log.Fatalln(err)
	}

	// Send Packet Burst
	for  i := 0; i < 10; i++ {
		_, err = conn.Write([]byte(strconv.Itoa(i)))
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println("Message sent: "+strconv.Itoa(i))
	}

	// Receive Packet Burst?
	for i := 0; i < 10; i++ 	{
		buffer := make([]byte, 1000)
		dataSize, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("connection closed")
			return
		}

		data := buffer[:dataSize]
		fmt.Println("received message: ", string(data))
	}
}

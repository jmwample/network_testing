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
	sent, k := 0, 15
	if err != nil {
		log.Fatalln(err)
	}

	// Send Packet Burst
	for  i := 0; i < k; i++ {
		_, err = conn.Write([]byte(strconv.Itoa(i)))
		if err != nil {
			log.Fatalln(err)
		}
		sent += len([]byte(strconv.Itoa(i)))
		fmt.Println("Message sent: "+strconv.Itoa(i))
	}

	// Receive Packet Burst?
	for i := 0; i < k; i++ 	{
		buffer := make([]byte, 1000)
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("connection closed")
			return
		}
		data := buffer[:n]
		fmt.Println("received message: ", string(data))

		sent -= n
		if sent == 0 {
			fmt.Println("connection closed")
			return
		}
	}
}

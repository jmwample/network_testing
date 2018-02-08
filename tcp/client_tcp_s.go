package tcp
// Serial
// 		TCP Packets are sent one by one. waiting for a response then
// 		sending more.


import (
	"fmt"
	"log"
	"net"
	"strconv"
)

const PortTCP = 16667

func main() {
	k := 12
	conn, err := net.Dial("tcp", fmt.Sprintf("localhost:%d", PortTCP))
	if err != nil {
		log.Fatalln(err)
	}

	// Send & Receive Packets
	for  i := 0; i < k; i++ {
		// SEND
		_, err = conn.Write([]byte(strconv.Itoa(i)))
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println("Message sent: "+strconv.Itoa(i))

		// RECEIVE
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

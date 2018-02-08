package main

import (
	"net"
	"fmt"
	"sync"
	"io"
)

const PORT = 16667


func main() {
	var wg sync.WaitGroup

	ln, err := net.Listen("tcp", fmt.Sprintf(":%d", PORT))
	if err != nil {
		fmt.Print(err)
		return
	}
	defer ln.Close()

	fmt.Printf("[SVR] TCP echo server listening on port %d\n", PORT)


	conn := receiveConn(ln)
	for {
		wg.Add(1)
		go echo(<-conn, &wg)

		wg.Wait()
	}
}


func receiveConn(listener net.Listener) chan net.Conn{
	ch := make(chan net.Conn)
	go func() {
		for {
			client, err := listener.Accept()
			if client == nil {
				fmt.Printf("[ERR[ Couldn't Accept Connection: " + err.Error())
				continue
			}
			fmt.Printf("[NEW] <--> %v\n", client.RemoteAddr())
			ch <- client
		}
	}()
	return ch
}


func echo(client net.Conn,  wg *sync.WaitGroup) {
	defer client.Close()
	defer wg.Done()


	for {
		msg := make([]byte, 1000)

		n, err := client.Read(msg)
		if err == io.EOF {
			fmt.Printf("[EOF] %v (%d bytes ignored)\n", client.RemoteAddr(), n)
			return
		} else if err != nil {
			fmt.Printf("[ERR] read\n%v", err.Error())
			return
		}

		fmt.Printf("[SVR] <--  %v\t%v\n", client.RemoteAddr(), msg[:n])

		n, err = client.Write([]byte(msg[:n]))
		if err != nil {
			fmt.Printf("[ERR] write\n")
			fmt.Print(err)
			return
		}

		fmt.Printf("[SVR] --> %v\n", client.RemoteAddr())

	}
}
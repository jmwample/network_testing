package main

import (
    "net"
    "fmt"
    "sync"
    "io"
)


func echo_srv(c net.Conn, wg *sync.WaitGroup) {
    defer c.Close()
    defer wg.Done()

    for {
        msg := make([]byte, 1000)

        n, err := c.Read(msg)
        if err == io.EOF {
            fmt.Printf("SERVER: received EOF (%d bytes ignored)\n", n)
            return
        } else  if err != nil {
            fmt.Printf("ERROR: read\n")
            fmt.Print(err)
            return
        }
        fmt.Printf("SERVER: received %v bytes\n", n)

        n, err = c.Write(msg[:n])
        if err != nil {
            fmt.Printf("ERROR: write\n")
            fmt.Print(err)
            return
        }
        fmt.Printf("SERVER: sent %v bytes\n", n)
    }
}

func main() {
    var wg sync.WaitGroup
    var echo_port = 16667

    ln, err := net.Listen("tcp", fmt.Sprintf(":%d", echo_port))

    fmt.Printf("TCP echo server listening on port %d\n", echo_port)
    if err != nil {
            fmt.Print(err)
            return
    }
    defer ln.Close()

    for {
        conn, err := ln.Accept()
        if err != nil {
                fmt.Print(err)
                return
        }
        wg.Add(1)
        go echo_srv(conn, &wg)

        wg.Wait()
    }

}

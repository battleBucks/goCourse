package main

import (
        "io"
        "net"
        "log"
)

func main() {
        listener, err := net.Listen("tcp", ":8080")

        if err != nil {
                log.Panic(err)
        }
        defer listener.Close()

        for {
                conn, err := listener.Accept()
                if err != nil {
                        log.Panic(err)
                }
                io.WriteString(conn, "\nI see you connected\n")
        }
}

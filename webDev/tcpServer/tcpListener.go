package main

import (
        "fmt"
        "io"
        "log"
        "net"
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
                io.WriteString(conn, "\nHello from TCP server\n")
                fmt.Fprintln(conn, "How is your day?")
                fmt.Fprintf(conn, "%v", "Well, I hope!\n")

                conn.Close()
        }
}
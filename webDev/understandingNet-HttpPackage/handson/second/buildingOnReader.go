package main

import (
        "io"
        "bufio"
        "net"
        "log"
        "fmt"
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
                go handle(conn)
        }
}

func handle(conn net.Conn) {
        scanner := bufio.NewScanner(conn)
        for scanner.Scan() {
                ln := scanner.Text()
                fmt.Println(ln)
                if ln == "" {
                        fmt.Println("THIS IS SOMETHING SOMETHING")
                        break
                }
        }
        defer conn.Close()
        fmt.Println("Code got here")
        io.WriteString(conn, "\nI see you connected\n")
}

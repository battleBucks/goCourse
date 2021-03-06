package main

import (
        "fmt"
        "bufio"
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
                go handle(conn)
        }
}

func handle(conn net.Conn) {
        scanner := bufio.NewScanner(conn)
        for scanner.Scan() {
                ln := scanner.Text()
                fmt.Println(ln)
                fmt.Fprintf(conn, "I heard you say: %s\n", ln)
        }
        defer conn.Close()

        fmt.Println("Do we ever get here?")
}

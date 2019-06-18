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
        defer conn.Close()
        scanner := bufio.NewScanner(conn)
        for scanner.Scan() {
                ln := scanner.Text()
                fmt.Println(ln)
                if ln == "" {
                        fmt.Println("THIS IS SOMETHING SOMETHING")
                        break
                }
        }
        body := "Check out this sweet response body payload"
        io.WriteString(conn, "HTTP/1.1 200 OK\r\n")
        fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
        fmt.Fprint(conn, "Content-Type: text/plain\r\n")
        io.WriteString(conn, "\r\n")
        io.WriteString(conn, body)
}

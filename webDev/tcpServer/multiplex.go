package main

import (
        "fmt"
        "log"
        "net"
        "strings"
        "bufio"
)

func main() {
        li, err := net.Listen("tcp", ":8080")
        if err != nil {
                log.Fatalln(err.Error())
        }
        defer li.Close()

        for {
                conn, err := li.Accept()
                if err != nil {
                        log.Fatalln(err.Error())
                }
                go handle(conn)
        }
}

func handle(conn net.Conn) {
        defer conn.Close()
        request(conn)
}

func request(conn net.Conn){
        i := 0
        scanner := bufio.NewScanner(conn)
        for scanner.Scan() {
                ln := scanner.Text()
                fmt.Println(ln)
                if i == 0 {
                        mux(conn, ln)
                }
                if ln == "" {
                        //headers are done
                        break
                }
                i++
        }
}
func mux(conn net.Conn, ln string){


        // request line
        method := strings.Fields(ln)[0]
        uri := strings.Fields(ln)[1]
        fmt.Println("***METHOD", method)
        fmt.Println("***URI", uri)

        //multiplexer
        if method == "GET" {
                if uri == "/" {
                        pages(conn, "index")
                }
                if uri == "/about" {
                        pages(conn, "about")
                }
                if uri == "/contact" {
                        pages(conn, "contact")
                }
                if uri == "/apply" {
                        pages(conn, "apply")
                }
        } else if method == "POST" && uri == "/apply" {
                applyProcess(conn)
        }
}

func pages(conn net.Conn, page string) {
        caps := strings.ToUpper(page)
        body := `<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="UTF-8">
	<title>Joe</title>
</head>
<body>
<strong>` + caps + `</strong><br>
<a href="/">index</a><br>
<a href="/about">about</a><br>
<a href="/contact">contact</a><br>
<a href="/apply">apply</a><br>
</body>
</html>`
        fmt.Fprint(conn, "HTTP/1.1 200 OK \r\n")
        fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
        fmt.Fprint(conn, "Content-Type: text/html\r\n")
        fmt.Fprint(conn, "\r\n")
        fmt.Fprint(conn, body)
}

func applyProcess(conn net.Conn) {
        body := `<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="UTF-8">
	<title>Joe</title>
</head>
<body>
<strong>APPLY PROCESS</strong><br>
<a href="/">index</a><br>
<a href="/about">about</a><br>
<a href="/contact">contact</a><br>
<a href="/apply">apply</a><br>
</body>
</html>`
        fmt.Fprint(conn, "HTTP/1.1 200 OK \r\n")
        fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
        fmt.Fprint(conn, "Content-Type: text/html\r\n")
        fmt.Fprint(conn, "\r\n")
        fmt.Fprint(conn, body)
}

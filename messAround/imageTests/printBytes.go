package main

import (
	"fmt"
	"bufio"
	"bytes"
//	"net/http"
	"os"
)

func main() {
//	filename := "image.jpg"
	filename := "/home/jbarrows/gopath/src/goCourse/messAround/imageTests/tester.jpg"
	file, err := os.Open(filename)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer file.Close()

	fileInfo, _ := file.Stat()
	var size int64 = fileInfo.Size()
	im_bytes := make([]byte, size)

	// read file into bytes
	buffer := bufio.NewReader(file)
	_, err = buffer.Read(im_bytes)
        buf := bytes.NewBuffer(im_bytes)
        fmt.Println(buf)
        return

	//filetype := http.DetectContentType(bytes)

        index := 0

        for int64(index) < 5000 {
        //for int64(index) < size - 1 {
                index += 1
                fmt.Printf("%c",im_bytes[index])
        }
}

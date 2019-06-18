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
	data := make([]byte, 0, size)
	temp_data := make([]byte, 0, size)

	// read file into bytes
	buffer := bufio.NewReader(file)
	_, err = buffer.Read(im_bytes)

	//filetype := http.DetectContentType(bytes)

        index := 0

        for int64(index + 9) < size {
                if bytes.Equal(im_bytes[index:index+10], []byte("<x:xmpmeta")) {
                        break
                }
                index += 1
        }
        for int64(index + 11) < size {
                if bytes.Equal(im_bytes[index:index+12], []byte("</x:xmpmeta>")) {
                        data = append(temp_data, im_bytes[index:index+1]...)
                        break
                }
                index += 1
        }
        if string(data) == "" {
                fmt.Println("Nothing to see here...")
        } else {
                fmt.Println(string(data))
        }
        //fmt.Println(data)
        //fmt.Println(temp_data)
        //fmt.Println(string(im_bytes[:7000]))
        fmt.Printf("%+ s",im_bytes[:7000])
}

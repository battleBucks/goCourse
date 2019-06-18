package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	fileToOpen := "nonExistant.txt"
	_, err := os.Open(fileToOpen)

	if err != nil {
		fmt.Printf("File: %s does not exist.\n%s\n", fileToOpen, err)
		log.Printf("File: %s does not exist.\n%s\n", fileToOpen, err)
		log.Fatalf("File: %s does not exist.\n%s\n", fileToOpen, err)
	}
}

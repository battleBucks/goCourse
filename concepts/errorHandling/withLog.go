package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	fileToCreate := "log.txt"
	fileToOpen := "nonExistant.txt"

	f, err := os.Create(fileToCreate)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	log.SetOutput(f)

	f2, err := os.Open(fileToOpen)

	if err != nil {
		fmt.Printf("File: %s does not exist.\n%s\n", fileToOpen, err)
		log.Printf("File: %s does not exist.\n%s\n", fileToOpen, err)
		log.Fatalf("File: %s does not exist.\n%s\n", fileToOpen, err)
	}
	defer f2.Close()
	fmt.Println("loggy log? are you there?")
}

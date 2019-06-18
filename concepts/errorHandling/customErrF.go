package main

import (
        "log"
        "fmt"
)

func main() {

        _, err := myErrFunc()
        if err != nil {
                log.Fatalln(err)
        }

}

func myErrFunc()(float64, error) {
        return 0, fmt.Errorf("This is my error message")
}

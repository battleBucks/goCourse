package main

import (
        "errors"
        "log"
)

func main() {

        _, err := myErrFunc()
        if err != nil {
                log.Fatalln(err)
        }

}

func myErrFunc()(float64, error) {
        return 0,errors.New("This is my error message")
}

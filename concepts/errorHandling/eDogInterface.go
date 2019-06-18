package main

import (
        "fmt"
        "log"
)

type joeB struct {
        name string
        awesomeness bool
}

func (j joeB) Error() string {
        return fmt.Sprintf("A Joe B error occured: %v\t|\t%v",j.name, j.awesomeness)
}

func main() {

        _, err := myFunc()
        if err != nil {
                log.Println(err)
        }
}

func myFunc()(float64, error) {
        return 0, joeB{"Joe", true}
}

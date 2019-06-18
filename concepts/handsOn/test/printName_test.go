package main

import (
//        "fmt"
        "testing"
)

func TestGetName(t *testing.T) {
        name := getName()
        if name != "Joseph Barrows" {
                t.Errorf("incorrect name")
        }
}

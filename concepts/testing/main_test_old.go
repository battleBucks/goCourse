package main

import (
        "testing"
)

func OldTestMySum(t *testing.T) {

        if mySum(5,10) != 15 {
                t.Errorf("Something happened with math")
        }
}

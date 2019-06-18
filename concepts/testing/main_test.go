package main

import (
        "testing"
)

func TestMySum(t *testing.T) {

        type test struct {
                data []int
                answer int
        }

        tests := []test{
                test{[]int{20,22}, 42},
                test{[]int{17,13}, 30},
                test{[]int{1,2,3,4,5}, 15},
                test{[]int{7}, 7},
                test{[]int{-2, 0, 3}, 1},
        }

        for _, v := range tests {
                sum := mySum(v.data...)
                if sum != v.answer {
                        t.Errorf("Something happened with math... Expected %d and got %d\n", v.answer, sum)
                }
        }
}

// Package sumIt sums it up
package sumIt

// MySum adds an unlimited number of ints and returns that as an int
func MySum(numbers ...int) int {
        sum := 0

        for _,val := range numbers {
                sum += val
        }

        return sum
}

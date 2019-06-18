package dog

import (
        "testing"
        "fmt"
)

func TestYears(t *testing.T) {
        humanYears := 3
        age := Years(humanYears)
        ans := humanYears * 7
        if age != ans {
                t.Errorf("Expecting %d but got %d.\n", ans, age)
        }
}
func TestYearsTwo(t *testing.T) {
        humanYears := 3
        age := YearsTwo(3)
        ans := humanYears * 7
        if age != ans {
                t.Errorf("Expecting %d but got %d.\n", ans, age)
        }
}
func ExampleYears() {
        fmt.Println(Years(10))
        // Output:
        // 70
}
func ExampleYearsTwo() {
        fmt.Println(YearsTwo(10))
        // Output:
        // 70
}
func BenchmarkYears(b *testing.B) {
        for i := 0; i < b.N; i++ {
                Years(10)
        }
}

func BenchmarkYearsTwo(b *testing.B) {
        for i := 0; i < b.N; i++ {
                YearsTwo(10)
        }
}

package sumIt

import (
        "fmt"
        "testing"
)

func ExampleMySum() {
        fmt.Println(MySum(2,3))
        // Output:
        // 5
}

func BenchmarkMySum(b *testing.B) {
        for i := 0; i < b.N; i++ {
                MySum((i+i), (i*i))
        }
}

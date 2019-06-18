package word

import (
        "testing"
        "fmt"
)

func ExampleCount() {
        fmt.Println(Count("sample sample string"))
        // Output:
        // 3
}
func BenchmarkUseCount(b *testing.B) {
        for i := 0; i < b.N; i++ {
                _ = UseCount("sample sample string string string")
        }
}

func BenchmarkCount(b *testing.B) {
        for i := 0; i < b.N; i++ {
                _ = Count("sample sample string string string")
        }
}

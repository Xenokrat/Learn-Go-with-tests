package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
    repreated := Repeat("a", 7)
    expected := "aaaaaaa"

    if repreated != expected {
        t.Errorf("expected %q, but got %q", expected, repreated)
    }
}

func BenchmarkRepeat(b *testing.B) {
    // ...setup...
    for b.Loop() {
        // code to measure
        Repeat("a", 7) 
    }
    // ...cleanup...
}

func ExampleRepeat() {
    repeated := Repeat("a", 5)
    fmt.Println(repeated)
    // Output: aaaaa
}

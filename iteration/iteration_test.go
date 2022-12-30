package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 3)
	expected := "aaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

// go test -bench=. -benchtime 100x
func BenchmarkRepeat(b *testing.B) {
	for i:= 0; i < b.N; i++ {
		Repeat("a", b.N)
	}
}

func ExampleRepeat() {
	repeated := Repeat("ba", 2)
	fmt.Println(repeated)
	// Output: baba
}

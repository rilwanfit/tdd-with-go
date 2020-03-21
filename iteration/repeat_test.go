package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("r", 5)
	expected := "rrrrr"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("r", 5)
	}
}

func ExampleRepeat() {
	repeated := Repeat("r", 5)
	fmt.Println(repeated)
	// Output: "rrrrr"
}

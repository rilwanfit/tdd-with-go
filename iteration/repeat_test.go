package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("r")
	expected := "rrrrr"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("r")
	}
}

func ExampleRepeat() {
	repeated := Repeat("r")
	fmt.Println(repeated)
	// Output: "rrrrr"
}

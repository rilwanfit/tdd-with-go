package iteration

import "testing"

func TestRepeat(t *testing.T) {
	repeated := Repeat("r")
	expected := "rrrrr"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

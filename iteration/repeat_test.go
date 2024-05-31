package iteration

import (
	"fmt"
	"testing"
)

func ExampleRepeat() {
	repeat := Repeat("a")
	fmt.Println(repeat)
	// Output: aaaaa
}
func TestRepeat(t *testing.T) {
	repeated := Repeat("a")
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected %q nut got %q", expected, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a")
	}
}

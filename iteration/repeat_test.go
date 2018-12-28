package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	actual := Repeat("a", 5)
	expected := "aaaaa"

	if actual != expected {
		t.Errorf("expected: %s; actual: %s", expected, actual)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

func ExampleRepeat() {
	repeated := Repeat("a", 5)
	fmt.Println(repeated)
	// Output: aaaaa
}

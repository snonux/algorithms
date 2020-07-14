package sort

import (
	"testing"
)

func BenchmarkSelection(b *testing.B) {
	benchmark("Selection", 10, Selection, b)
	benchmark("Selection", 100, Selection, b)
	benchmark("Selection", 1000, Selection, b)
}

func TestSelection(t *testing.T) {
	for i := 1; i <= 1000; i *= 10 {
		test("Selection", i, Selection, t)
	}
}

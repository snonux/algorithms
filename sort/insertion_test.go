package sort

import (
	"testing"
)

func BenchmarkInsertion(b *testing.B) {
	benchmark("Insertion", 10, Insertion, b)
	benchmark("Insertion", 100, Insertion, b)
	benchmark("Insertion", 1000, Insertion, b)
}

func TestInsertion(t *testing.T) {
	for i := 1; i <= 1000; i *= 10 {
		test("Insertion", i, Insertion, t)
	}
}

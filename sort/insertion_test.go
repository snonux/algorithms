package sort

import (
	"algorithms/ds"
	"testing"
)

func BenchmarkInsertion(b *testing.B) {
	benchmark("Insertion", 10, Insertion, b)
	benchmark("Insertion", 100, Insertion, b)
	benchmark("Insertion", 1000, Insertion, b)
}

func TestInsertion1000(t *testing.T) {
	a := ds.RandomIntegers(1000, 1000)
	a = Insertion(a)
	if !Sorted(a) {
		t.Errorf("Array not sorted: %v", a)
	}
}

package sort

import (
	"algorithms/ds"
	"testing"
)

func BenchmarkSelection(b *testing.B) {
	benchmark("Selection", 10, Selection, b)
	benchmark("Selection", 100, Selection, b)
	benchmark("Selection", 1000, Selection, b)
}

func TestSelection1000(t *testing.T) {
	a := ds.RandomIntegers(1000, 1000)
	a = Selection(a)
	if !Sorted(a) {
		t.Errorf("Array not sorted: %v", a)
	}
}

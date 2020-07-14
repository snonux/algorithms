package sort

import (
	"algorithms/ds"
	"testing"
)

func BenchmarkInsertion1000(b *testing.B) {
	a := ds.RandomIntegers(1000, 1000)
	b.ResetTimer()
	Insertion(a)
}

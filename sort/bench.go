package sort

import (
	"algorithms/ds"
	"fmt"
	"testing"
)

// Avoid compiler optimizations
var benchResult []ds.Comparer

func benchmark(name string, length int, sort func([]ds.Comparer) []ds.Comparer, b *testing.B) {
	cb := func(b *testing.B) {
		a := ds.RandomIntegers(length, length)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			benchResult = sort(a)
		}
	}
	b.Run(fmt.Sprintf("Benchmark%s%d", name, length), cb)
}

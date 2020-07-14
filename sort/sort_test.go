package sort

import (
	"algorithms/ds"
	"fmt"
	"testing"
)

// Store results here to avoid compiler optimizations
var benchResult []ds.Comparer

const maxLength int = 10000

type sortAlgorithm func([]ds.Comparer) []ds.Comparer

func TestSort(t *testing.T) {
	for i := 1; i <= maxLength; i *= 10 {
		test("Selection", i, Selection, t)
		test("Insertion", i, Insertion, t)
	}
}

func BenchmarkSort(b *testing.B) {
	for i := 1; i <= maxLength; i *= 10 {
		benchmark("Insertion", i, Insertion, b)
	}
	for i := 1; i <= maxLength; i *= 10 {
		benchmark("Selection", i, Selection, b)
	}
}

func test(name string, length int, sort sortAlgorithm, t *testing.T) {
	cb := func(t *testing.T) {
		t.Parallel()
		a := ds.RandomIntegers(length, length)
		a = sort(a)
		if !Sorted(a) {
			t.Errorf("Array not sorted: %v", a)
		}
	}
	t.Run(fmt.Sprintf("Test%s%d", name, length), cb)
}

func benchmark(name string, length int, sort sortAlgorithm, b *testing.B) {
	cb := func(b *testing.B) {
		a := ds.RandomIntegers(length, length)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			benchResult = sort(a)
		}
	}
	b.Run(fmt.Sprintf("Benchmark%s%d", name, length), cb)
}

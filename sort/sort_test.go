package sort

import (
	"algorithms/ds"
	"fmt"
	"testing"
)

// Store results here to avoid compiler optimizations
var benchResult ds.ArrayList
var benchResultInt []int

const maxLength int = 10000

type sortAlgorithm func(ds.ArrayList) ds.ArrayList
type sortAlgorithmInt func([]int) []int

func TestSelectionSort(t *testing.T) {
	for i := 1; i <= maxLength; i *= 10 {
		test(Selection, i, t)
	}
}

func TestInsertionSort(t *testing.T) {
	for i := 1; i <= maxLength; i *= 10 {
		test(Insertion, i, t)
	}
}

func TestShellSort(t *testing.T) {
	for i := 1; i <= maxLength; i *= 10 {
		test(Shell, i, t)
	}
}

func TestMergeSort(t *testing.T) {
	for i := 1; i <= maxLength; i *= 10 {
		test(Merge, i, t)
	}
}

func TestShuffleSort(t *testing.T) {
	for i := 10; i <= maxLength; i *= 10 {
		testShuffle(Shuffle, i, t)
	}
}

func BenchmarkInsertionSort(b *testing.B) {
	for i := 1; i <= maxLength; i *= 10 {
		benchmark(Insertion, i, b)
	}
}

func BenchmarkSelectionSort(b *testing.B) {
	for i := 1; i <= maxLength; i *= 10 {
		benchmark(Selection, i, b)
	}
}

func BenchmarkShellSort(b *testing.B) {
	for i := 1; i <= maxLength; i *= 10 {
		benchmark(Shell, i, b)
	}
}

func BenchmarkMergeSort(b *testing.B) {
	for i := 1; i <= maxLength; i *= 10 {
		benchmark(Merge, i, b)
	}
}

func BenchmarkShuffleSort(b *testing.B) {
	for i := 1; i <= maxLength; i *= 10 {
		benchmark(Shuffle, i, b)
	}
}

func test(sort sortAlgorithm, length int, t *testing.T) {
	cb := func(t *testing.T) {
		t.Parallel()
		a := makeIntegers(length, length)
		a = sort(a)
		if !a.Sorted() {
			t.Errorf("Array not sorted: %v", a)
		}
	}
	t.Run(fmt.Sprintf("%d", length), cb)
}

func testShuffle(sort sortAlgorithm, length int, t *testing.T) {
	cb := func(t *testing.T) {
		t.Parallel()
		a := sort(ds.SortedIntegers(length))
		if a.Sorted() {
			t.Errorf("Array sorted: %v", a.FirstN(10))
		}
	}
	t.Run(fmt.Sprintf("%d", length), cb)
}

func benchmark(sort sortAlgorithm, length int, b *testing.B) {
	cb := func(b *testing.B) {
		a := makeIntegers(length, length)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			benchResult = sort(a)
		}
	}
	b.Run(fmt.Sprintf("%d", length), cb)
}

func makeIntegers(length, max int) ds.ArrayList {
	return ds.RandomIntegers(length, max)
}

package sort

import (
	"algorithms/ds"
	"fmt"
	"testing"
)

// Store results here to avoid compiler optimizations
var benchResult ds.ArrayList
var benchResultInt []int

const minLength int = 1
const maxLength int = 1000000
const maxSlowLength int = 100000

var arrayListCache map[string]ds.ArrayList

type sortAlgorithm func(ds.ArrayList) ds.ArrayList
type sortAlgorithmInt func([]int) []int

func TestSelectionSort(t *testing.T) {
	for i := minLength; i <= maxSlowLength; i *= 10 {
		test(Selection, i, t)
	}
}

func TestInsertionSort(t *testing.T) {
	for i := minLength; i <= maxSlowLength; i *= 10 {
		test(Insertion, i, t)
	}
}

func TestShellSort(t *testing.T) {
	for i := minLength; i <= maxLength; i *= 10 {
		test(Shell, i, t)
	}
}

func TestMergeSort(t *testing.T) {
	for i := minLength; i <= maxLength; i *= 10 {
		test(Merge, i, t)
	}
}

func TestBottomUpMergeSort(t *testing.T) {
	t.Log("Parallel merge sort")
	for i := minLength; i <= maxLength; i *= 10 {
		test(BottomUpMerge, i, t)
	}
	test(BottomUpMerge, maxLength*2, t)
}

func TestParallelMergeSort(t *testing.T) {
	t.Log("Bottom-up merge sort")
	for i := minLength; i <= maxLength; i *= 10 {
		test(ParallelMerge, i, t)
	}
}

func TestQuickSort(t *testing.T) {
	for i := minLength; i <= maxLength; i *= 10 {
		test(Quick, i, t)
	}
}

func TestParallelQuickSort(t *testing.T) {
	for i := minLength; i <= maxLength; i *= 10 {
		test(ParallelQuick, i, t)
	}
}

func TestQuick3WaySort(t *testing.T) {
	for i := minLength; i <= maxLength; i *= 10 {
		test(Quick3Way, i, t)
	}
}

func TestShuffleSort(t *testing.T) {
	for i := 10; i <= maxLength; i *= 10 {
		testShuffle(Shuffle, i, t)
	}
}

func BenchmarkInsertionSort(b *testing.B) {
	for i := minLength; i <= maxSlowLength; i *= 10 {
		benchmark(Insertion, i, b)
	}
}

func BenchmarkSelectionSort(b *testing.B) {
	for i := minLength; i <= maxSlowLength; i *= 10 {
		benchmark(Selection, i, b)
	}
}

func BenchmarkShellSort(b *testing.B) {
	for i := minLength; i <= maxLength; i *= 10 {
		benchmark(Shell, i, b)
	}
}

func BenchmarkMergeSort(b *testing.B) {
	for i := minLength; i <= maxLength; i *= 10 {
		benchmark(Merge, i, b)
	}
}

func BenchmarkBottomUpMergeSort(b *testing.B) {
	for i := minLength; i <= maxLength; i *= 10 {
		benchmark(BottomUpMerge, i, b)
	}
}

func BenchmarkParallelMergeSort(b *testing.B) {
	for i := minLength; i <= maxLength; i *= 10 {
		benchmark(ParallelMerge, i, b)
	}
}

func BenchmarkQuickSort(b *testing.B) {
	for i := minLength; i <= maxLength; i *= 10 {
		benchmark(Quick, i, b)
	}
}

func BenchmarkParallelQuickSort(b *testing.B) {
	for i := minLength; i <= maxLength; i *= 10 {
		benchmark(ParallelQuick, i, b)
	}
}

func BenchmarkQuick3WaySort(b *testing.B) {
	for i := minLength; i <= maxLength; i *= 10 {
		benchmark(Quick3Way, i, b)
	}
}

/*
func BenchmarkShuffleSort(b *testing.B) {
	for i := minLength; i <= maxLength; i *= 10 {
		benchmark(Shuffle, i, b)
	}
}
*/

func test(sort sortAlgorithm, length int, t *testing.T) {
	cb := func(t *testing.T) {
		t.Parallel()
		a := makeRandomIntegers(length, -1)
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
		a := sort(ds.AscendingIntegers(length))
		if a.Sorted() {
			t.Errorf("Array sorted: %v", a.FirstN(10))
		}
	}
	t.Run(fmt.Sprintf("%d", length), cb)
}

func benchmark(sort sortAlgorithm, length int, b *testing.B) {
	a := makeRandomIntegers(length, -1)
	b.Run(fmt.Sprintf("random(%d)", length), func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			benchResult = sort(a)
		}
	})

	a = ds.AscendingIntegers(length)
	b.Run(fmt.Sprintf("ascending(%d)", length), func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			benchResult = sort(a)
		}
	})

	a = ds.DescendingIntegers(length)
	b.Run(fmt.Sprintf("descending(%d)", length), func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			benchResult = sort(a)
		}
	})
}

func makeRandomIntegers(length, max int) ds.ArrayList {
	// Use a cache to make sure that all all sorting algos use the same
	// random sequences.
	if arrayListCache == nil {
		arrayListCache = make(map[string]ds.ArrayList)
	}

	key := fmt.Sprintf("random(%d, %d)", length, max)
	if a, ok := arrayListCache[key]; ok {
		return a
	}
	a := ds.RandomIntegers(length, max)
	a = Shuffle(a)
	arrayListCache[key] = a
	return a
}

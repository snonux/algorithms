package sort

import (
	"fmt"
	"testing"

	"github.com/snonux/algorithms/ds"
)

// Store results here to avoid compiler optimizations
var benchResult ds.ArrayList

const minLength int = 1
const maxLength int = 1000000
const factor int = 100
const maxSlowLength int = 100000

var arrayListCache map[string]ds.ArrayList

type sortAlgorithm func(ds.ArrayList) ds.ArrayList
type sortAlgorithmInt func([]int) []int

func TestSelectionSort(t *testing.T) {
	for i := minLength; i <= maxSlowLength; i *= factor {
		test(Selection, i, t)
	}
}

func TestInsertionSort(t *testing.T) {
	for i := minLength; i <= maxSlowLength; i *= factor {
		test(Insertion, i, t)
	}
}

func TestShellSort(t *testing.T) {
	for i := minLength; i <= maxLength; i *= factor {
		test(Shell, i, t)
	}
}

func TestMergeSort(t *testing.T) {
	for i := minLength; i <= maxLength; i *= factor {
		test(Merge, i, t)
	}
}

func TestBottomUpMergeSort(t *testing.T) {
	t.Log("Parallel merge sort")
	for i := minLength; i <= maxLength; i *= factor {
		test(BottomUpMerge, i, t)
	}
	test(BottomUpMerge, maxLength*2, t)
}

func TestParallelMergeSort(t *testing.T) {
	t.Log("Bottom-up merge sort")
	for i := minLength; i <= maxLength; i *= factor {
		test(ParallelMerge, i, t)
	}
}

func TestQuickSort(t *testing.T) {
	for i := minLength; i <= maxLength; i *= factor {
		test(Quick, i, t)
	}
}

func TestParallelQuickSort(t *testing.T) {
	for i := minLength; i <= maxLength; i *= factor {
		test(ParallelQuick, i, t)
	}
}

func TestQuick3WaySort(t *testing.T) {
	for i := minLength; i <= maxLength; i *= factor {
		test(Quick3Way, i, t)
	}
}

func TestShuffleSort(t *testing.T) {
	for i := 10; i <= maxLength; i *= factor {
		testShuffle(Shuffle, i, t)
	}
}

func BenchmarkInsertionSort(b *testing.B) {
	for i := minLength; i <= maxSlowLength; i *= factor {
		benchmark(Insertion, i, b)
	}
}

func BenchmarkSelectionSort(b *testing.B) {
	for i := minLength; i <= maxSlowLength; i *= factor {
		benchmark(Selection, i, b)
	}
}

func BenchmarkShellSort(b *testing.B) {
	for i := minLength; i <= maxLength; i *= factor {
		benchmark(Shell, i, b)
	}
}

func BenchmarkMergeSort(b *testing.B) {
	for i := minLength; i <= maxLength; i *= factor {
		benchmark(Merge, i, b)
	}
}

func BenchmarkBottomUpMergeSort(b *testing.B) {
	for i := minLength; i <= maxLength; i *= factor {
		benchmark(BottomUpMerge, i, b)
	}
}

func BenchmarkParallelMergeSort(b *testing.B) {
	for i := minLength; i <= maxLength; i *= factor {
		benchmark(ParallelMerge, i, b)
	}
}

func BenchmarkQuickSort(b *testing.B) {
	for i := minLength; i <= maxLength; i *= factor {
		benchmark(Quick, i, b)
	}
}

func BenchmarkParallelQuickSort(b *testing.B) {
	for i := minLength; i <= maxLength; i *= factor {
		benchmark(ParallelQuick, i, b)
	}
}

func BenchmarkQuick3WaySort(b *testing.B) {
	for i := minLength; i <= maxLength; i *= factor {
		benchmark(Quick3Way, i, b)
	}
}

/*
func BenchmarkShuffleSort(b *testing.B) {
	for i := minLength; i <= maxLength; i *= factor {
		benchmark(Shuffle, i, b)
	}
}
*/

func test(sort sortAlgorithm, l int, t *testing.T) {
	cb := func(t *testing.T) {
		t.Parallel()
		a := ds.NewRandomArrayList(l, -1)
		a = sort(a)
		if !a.Sorted() {
			t.Errorf("Array not sorted: %v", a)
		}
	}
	t.Run(fmt.Sprintf("%d", l), cb)
}

func testShuffle(sort sortAlgorithm, l int, t *testing.T) {
	cb := func(t *testing.T) {
		t.Parallel()
		a := sort(ds.NewAscendingArrayList(l))
		if a.Sorted() {
			t.Errorf("Array sorted: %v", a.FirstN(10))
		}
	}
	t.Run(fmt.Sprintf("%d", l), cb)
}

func benchmark(sort sortAlgorithm, l int, b *testing.B) {
	a := ds.NewRandomArrayList(l, -1)
	aux := ds.NewArrayList(l)

	b.Run(fmt.Sprintf("random(%d)", l), func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			copy(aux, a)
			benchResult = sort(aux)
		}
	})

	a = ds.NewAscendingArrayList(l)
	b.Run(fmt.Sprintf("ascending(%d)", l), func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			copy(aux, a)
			benchResult = sort(aux)
		}
	})

	a = ds.NewDescendingArrayList(l)
	b.Run(fmt.Sprintf("descending(%d)", l), func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			copy(aux, a)
			benchResult = sort(aux)
		}
	})
}

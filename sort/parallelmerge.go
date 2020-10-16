package sort

import (
	"sync"

	"github.com/snonux/algorithms/ds"
)

func ParallelMerge(a ds.ArrayList) ds.ArrayList {
	parallelMerge(a, make(ds.ArrayList, len(a)))
	return a
}

func parallelMerge(a, aux ds.ArrayList) {
	l := len(a)

	if l < 1000 {
		mergeSort(a, aux)
		return
	}

	mi := l / 2
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		parallelMerge(a[0:mi], aux[0:mi])
		wg.Done()
	}()

	go func() {
		parallelMerge(a[mi:], aux[mi:])
		wg.Done()
	}()

	wg.Wait()
	merge(a, aux, 0, mi, l-1)
}

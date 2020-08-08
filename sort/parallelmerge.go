package sort

import (
	"algorithms/ds"
	"sync"
)

func ParallelMerge(a ds.ArrayList) ds.ArrayList {
	parallelMerge(a, make(ds.ArrayList, len(a)))
	return a
}

func parallelMerge(a, aux ds.ArrayList) {
	length := len(a)
	if length <= 1 {
		return
	}

	mi := length / 2
	if length < 1000 {
		mergeSort(a[0:mi], aux[0:mi])
		mergeSort(a[mi:], aux[mi:])
		merge(a, aux, 0, mi, length-1)
		return
	}

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
	merge(a, aux, 0, mi, length-1)
	return
}

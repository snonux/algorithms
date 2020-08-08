package sort

import (
	"algorithms/ds"
	"sync"
)

func ParallelMerge(a ds.ArrayList) ds.ArrayList {
	aux := make(ds.ArrayList, len(a))
	parallelMerge(a, aux, 0, len(a)-1)

	return a
}

func parallelMerge(a, aux ds.ArrayList, lo, hi int) {
	mid := lo + (hi-lo)/2
	defer merge(a, aux, lo, mid, hi)

	if hi-lo <= 1000 {
		//mergeSort(a, aux, lo, mid)
		//mergeSort(a, aux, mid+1, hi)
		return
	}

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		parallelMerge(a, aux, lo, mid)
		wg.Done()
	}()
	go func() {
		parallelMerge(a, aux, mid+1, hi)
		wg.Done()
	}()
	wg.Wait()
}

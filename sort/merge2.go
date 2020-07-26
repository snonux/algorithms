package sort

import (
	"algorithms/ds"
	"sync"
)

// Merge2 is a parallelized version of Merge
func Merge2(a ds.ArrayList) ds.ArrayList {
	aux := make(ds.ArrayList, len(a))
	mergeSort2(a, aux, 0, len(a)-1)

	return a
}

func mergeSort2(a, aux ds.ArrayList, lo, hi int) {
	mid := lo + (hi-lo)/2
	defer merge(a, aux, lo, mid, hi)

	if hi-lo <= 1000 {
		mergeSort(a, aux, lo, mid)
		mergeSort(a, aux, mid+1, hi)
		return
	}

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		mergeSort2(a, aux, lo, mid)
		wg.Done()
	}()
	go func() {
		mergeSort2(a, aux, mid+1, hi)
		wg.Done()
	}()
	wg.Wait()
}

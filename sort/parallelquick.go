package sort

import (
	"algorithms/ds"
	"sync"
)

func ParallelQuick(a ds.ArrayList) ds.ArrayList {
	//Shuffle(a)
	parallelQuick(a)
	return a
}

func parallelQuick(a ds.ArrayList) {
	l := len(a)

	if l < 1000 {
		quick(a)
		return
	}

	j := quickPartition(a)
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		parallelQuick(a[0:j])
		wg.Done()
	}()
	go func() {
		parallelQuick(a[j+1:])
		wg.Done()
	}()

	wg.Wait()
}

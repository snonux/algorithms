package sort

import (
	"algorithms/ds"
	"sync"
)

func ParallelQuick(a ds.ArrayList) ds.ArrayList {
	Shuffle(a)
	parallelQuick(a)
	return a
}

func parallelQuick(a ds.ArrayList) {
	length := len(a)
	if length <= 10 {
		Insertion(a)
		return
	}

	j := quickPartition(a)

	if length >= 1000 {
		var wg sync.WaitGroup
		wg.Add(2)
		defer wg.Wait()

		go func() {
			parallelQuick(a[0:j])
			wg.Done()
		}()
		go func() {
			parallelQuick(a[j+1:])
			wg.Done()
		}()
		return
	}

	parallelQuick(a[0:j])
	parallelQuick(a[j+1:])
}

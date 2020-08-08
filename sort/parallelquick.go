package sort

import (
	"algorithms/ds"
	"sync"
)

func ParallelQuick(a ds.ArrayList) ds.ArrayList {
	Shuffle(a)
	parallelQuick(a, 0, len(a)-1)
	return a
}

func parallelQuick(a ds.ArrayList, lo, hi int) {
	if hi <= lo+10 {
		insertion(a, lo, hi)
		return
	}

	j := quickPartition(a, lo, hi)

	if hi <= lo+1000 {
		var wg sync.WaitGroup
		wg.Add(2)
		defer wg.Wait()

		go func() {
			parallelQuick(a, lo, j-1)
			wg.Done()
		}()
		go func() {
			parallelQuick(a, j+1, hi)
			wg.Done()
		}()
		return
	}

	parallelQuick(a, lo, j-1)
	parallelQuick(a, j+1, hi)
}

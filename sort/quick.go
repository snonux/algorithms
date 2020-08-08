package sort

import (
	"algorithms/ds"
)

func Quick(a ds.ArrayList) ds.ArrayList {
	Shuffle(a)
	quick(a)
	return a
}

func quick(a ds.ArrayList) {
	if len(a) <= 10 {
		Insertion(a)
		return
	}

	j := quickPartition(a)
	quick(a[0:j])
	quick(a[j+1:])
}

func quickPartition(a ds.ArrayList) int {
	i := 0      // Left scan index
	j := len(a) // Right scan index

	Insertion(a[0:3])
	a.Swap(0, 1)
	v := a[0] // Partitioning item

	for {
		for i++; a[i] < v; i++ {
		}

		for j--; v < a[j]; j-- {
		}

		// Check scan complete
		if i >= j {
			break
		}

		a.Swap(i, j)
	}

	// Put partitioning item v into a[j]
	a.Swap(0, j)

	// whith a[lo..j-1 <= a[j] <= a[j+1..hi]
	return j
}

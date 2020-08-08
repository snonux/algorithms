package sort

import (
	"algorithms/ds"
)

func Quick(a ds.ArrayList) ds.ArrayList {
	Shuffle(a)
	quick(a, 0, len(a)-1)
	return a
}

func quick(a ds.ArrayList, lo, hi int) {
	if hi <= lo+10 {
		insertion(a, lo, hi)
		return
	}

	j := quickPartition(a, lo, hi)

	quick(a, lo, j-1)
	quick(a, j+1, hi)
}

func quickPartition(a ds.ArrayList, lo, hi int) int {
	i := lo     // Left scan index
	j := hi + 1 // Right scan index

	insertion(a, lo, lo+2)
	a.Swap(lo, lo+1)
	v := a[lo] // Partitioning item

	for {
		for i++; a[i] < v; i++ {
			if i == hi {
				break
			}
		}

		for j--; v < a[j]; j-- {
			if j == lo {
				break
			}
		}

		// Check scan complete
		if i >= j {
			break
		}

		a.Swap(i, j)
	}

	// Put partitioning item v into a[j]
	a.Swap(lo, j)

	// whith a[lo..j-1 <= a[j] <= a[j+1..hi]
	return j
}

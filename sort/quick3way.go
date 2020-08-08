package sort

import (
	"algorithms/ds"
)

// Quick3Way uses a 3-way partitioning so it is more efficient dealing with duplicates
func Quick3Way(a ds.ArrayList) ds.ArrayList {
	Shuffle(a)
	quick3Way(a, 0, len(a)-1)
	return a
}

func quick3Way(a ds.ArrayList, lo, hi int) {
	if hi <= lo+10 {
		insertion(a, lo, hi)
		return
	}

	lt := lo    // Lower than
	i := lo + 1 // lt..i contain duplicates
	gt := hi    // Greater than

	insertion(a, lo, lo+2)
	a.Swap(lo, lo+1)
	v := a[lo] // Partitioning item

	for i <= gt {
		switch {
		case a[i] < v:
			a.Swap(lt, i)
			lt++
			i++
		case a[i] > v:
			a.Swap(i, gt)
			gt--
		default:
			// Duplicate
			i++
		}
	}
	// Now a[lo..lt-1] < v = a[lt..gt] < a[gt+1..hi]

	quick3Way(a, lo, lt-1)
	quick3Way(a, gt+1, hi)
}

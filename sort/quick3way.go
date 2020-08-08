package sort

import (
	"algorithms/ds"
)

// Quick3Way uses a 3-way partitioning so it is more efficient dealing with duplicates
func Quick3Way(a ds.ArrayList) ds.ArrayList {
	Shuffle(a)
	quick3Way(a)
	return a
}

func quick3Way(a ds.ArrayList) {
	l := len(a)
	if l <= 10 {
		Insertion(a)
		return
	}

	lt := 0     // Lower than
	i := 1      // lt..i contain duplicates
	gt := l - 1 // Greater than

	a.Swap(0, median(a, l))
	v := a[0]

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

	quick3Way(a[0:lt])
	quick3Way(a[gt+1:])
}

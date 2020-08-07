package sort

import (
	"algorithms/ds"
)

// Quick2 uses a 3-way partitioning so it is more efficient
// dealing with duplicates
func Quick2(a ds.ArrayList) ds.ArrayList {
	a = Shuffle(a)
	quick2(a, 0, len(a)-1)
	return a
}

func quick2(a ds.ArrayList, lo, hi int) {
	if hi <= lo {
		return
	}

	lt := lo    // Lower than
	i := lo + 1 // lt..i contain duplicates
	gt := hi    // Greater than
	v := a[lo]  // Partitioning item

	for i <= gt {
		switch a[i].Compare(v) {
		case -1:
			a.Swap(lt, i)
			lt++
			i++
		case 1:
			a.Swap(i, gt)
			gt--
		default:
			// Duplicate
			i++
		}
	}
	// Now a[lo..lt-1] < v = a[lt..gt] < a[gt+1..hi]

	quick2(a, lo, lt-1)
	quick2(a, gt+1, hi)
}

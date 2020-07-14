package sort

import "algorithms/ds"

func Sorted(a []ds.Comparer) bool {
	for i := len(a) - 1; i > 0; i-- {
		if a[i].LowerThan(a[i-1]) {
			return false
		}
	}
	return true
}

package sort

import (
	"algorithms/ds"
)

// Merge3 is the bottom up version of merge sort.
func Merge3(a ds.ArrayList) ds.ArrayList {
	length := len(a)
	aux := make(ds.ArrayList, length)

	for sz := 1; sz < length; sz = sz + sz {
		for lo := 0; lo < length-sz; lo += sz + sz {
			merge(a, aux, lo, lo+sz-1, min(lo+sz+sz-1, length-1))
		}
	}

	return a
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

package sort

import (
	"github.com/snonux/algorithms/ds"
)

func BottomUpMerge(a ds.ArrayList) ds.ArrayList {
	l := len(a)
	aux := make(ds.ArrayList, l)

	for sz := 1; sz < l; sz = sz + sz {
		for lo := 0; lo < l-sz; lo += sz + sz {
			merge(a, aux, lo, lo+sz, min(lo+sz+sz-1, l-1))
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

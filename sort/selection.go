package sort

import (
	"github.com/snonux/algorithms/ds"
)

func Selection(a ds.ArrayList) ds.ArrayList {
	l := len(a)

	for i := 0; i < l; i++ {
		min := i
		for j := i + 1; j < l; j++ {
			if a[min] > a[j] {
				min = j
			}
		}
		if min == i {
			continue
		}
		a.Swap(i, min)
	}

	return a
}

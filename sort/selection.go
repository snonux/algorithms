package sort

import (
	"algorithms/ds"
)

func Selection(a []ds.Comparer) []ds.Comparer {
	length := len(a)
	for i := 0; i < length; i++ {
		max := i
		for j := i + 1; j < length; j++ {
			if a[max].HigherThan(a[j]) {
				max = j
			}
		}
		if max == i {
			continue
		}
		tmp := a[i]
		a[i] = a[max]
		a[max] = tmp
	}
	return a
}

package sort

import (
	"algorithms/ds"
)

func Insertion(a []ds.Comparer) []ds.Comparer {
	length := len(a)

	for i := 0; i < length; i++ {
		for j := i; j > 0; j-- {
			if a[j].HigherThan(a[j-1]) {
				break
			}
			tmp := a[j]
			a[j] = a[j-1]
			a[j-1] = tmp
		}
	}

	return a
}

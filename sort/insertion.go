package sort

import (
	"algorithms/ds"
)

func Insertion(a ds.ArrayList) ds.ArrayList {
	insertion(a, 0, len(a)-1)
	return a
}

func insertion(a ds.ArrayList, lo, hi int) {
	for i := lo; i <= hi; i++ {
		for j := i; j > 0; j-- {
			if a[j].Higher(a[j-1]) {
				break
			}
			a.Swap(j, j-1)
		}
	}
}

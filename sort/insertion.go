package sort

import (
	"github.com/snonux/algorithms/ds"
)

func Insertion(a ds.ArrayList) ds.ArrayList {
	for i, _ := range a {
		for j := i; j > 0; j-- {
			if a[j] > a[j-1] {
				break
			}
			a.Swap(j, j-1)
		}
	}

	return a
}

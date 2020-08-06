package sort

import (
	"algorithms/ds"
)

func Insertion(a ds.ArrayList) ds.ArrayList {
	length := len(a)

	for i := 0; i < length; i++ {
		for j := i; j > 0; j-- {
			if a[j].Higher(a[j-1]) {
				break
			}
			a.Swap(j, j-1)
		}
	}

	return a
}

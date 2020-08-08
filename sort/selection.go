package sort

import (
	"algorithms/ds"
)

func Selection(a ds.ArrayList) ds.ArrayList {
	length := len(a)

	for i := 0; i < length; i++ {
		min := i
		for j := i + 1; j < length; j++ {
			if a[min].Higher(a[j]) {
				min = j
			}
		}
		if min == i {
			continue
		}
		tmp := a[i]
		a[i] = a[min]
		a[min] = tmp
	}

	return a
}

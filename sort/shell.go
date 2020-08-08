package sort

import (
	"algorithms/ds"
)

func Shell(a ds.ArrayList) ds.ArrayList {
	l := len(a)
	// h-sort the array
	h := 1

	for h < l/3 {
		// 1, 4, 13, 40, 121, 364, 1093...
		h = 3*h + 1
	}

	for h >= 1 {
		for i := h; i < l; i++ {
			for j := i; j >= h; j -= h {
				if a[j-h] < a[j] {
					break
				}
				a.Swap(j, j-h)
			}
		}

		h /= 3
	}

	return a
}

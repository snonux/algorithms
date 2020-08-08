package sort

import (
	"algorithms/ds"
)

func Merge(a ds.ArrayList) ds.ArrayList {
	aux := make(ds.ArrayList, len(a))
	mergeSort(a, aux, 0, len(a)-1)

	return a
}

func mergeSort(a, aux ds.ArrayList, lo, hi int) {
	if lo >= hi {
		return
	}
	mid := lo + (hi-lo)/2

	mergeSort(a, aux, lo, mid)
	mergeSort(a, aux, mid+1, hi)
	merge(a, aux, lo, mid, hi)
}

func merge(a, aux ds.ArrayList, lo, mid, hi int) {
	for i := lo; i <= hi; i++ {
		aux[i] = a[i]
	}

	i := lo
	j := mid + 1

	for k := lo; k <= hi; k++ {
		switch {
		case i > mid:
			a[k] = aux[j]
			j++
		case j > hi:
			a[k] = aux[i]
			i++
		case aux[i] > aux[j]:
			a[k] = aux[j]
			j++
		default:
			a[k] = aux[i]
			i++
		}
	}
}

package sort

import (
	"github.com/snonux/algorithms/ds"
)

func Merge(a ds.ArrayList) ds.ArrayList {
	aux := make(ds.ArrayList, len(a))
	mergeSort(a, aux)

	return a
}

func mergeSort(a, aux ds.ArrayList) {
	l := len(a)
	if l <= 10 {
		Insertion(a)
		return
	}

	mi := l / 2
	mergeSort(a[0:mi], aux[0:mi])
	mergeSort(a[mi:], aux[mi:])
	merge(a, aux, 0, mi, l-1)
}

func merge(a, aux ds.ArrayList, lo, mi, hi int) {
	for i := lo; i <= hi; i++ {
		aux[i] = a[i]
	}

	i := lo
	j := mi

	for k := lo; k <= hi; k++ {
		switch {
		case i >= mi:
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

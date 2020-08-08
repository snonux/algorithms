package sort

import (
	"algorithms/ds"
	"math/rand"
)

func Quick(a ds.ArrayList) ds.ArrayList {
	quick(a)
	return a
}

func quick(a ds.ArrayList) {
	if len(a) <= 10 {
		Insertion(a)
		return
	}

	j := quickPartition(a)
	quick(a[0:j])
	quick(a[j+1:])
}

func quickPartition(a ds.ArrayList) int {
	l := len(a)
	i := 0 // Left scan index
	j := l // Right scan index
	hi := j - 1

	a.Swap(0, median(a, l))
	v := a[0]

	for {
		for i++; a[i] < v && i < hi; i++ {
		}

		for j--; v < a[j] && j > 0; j-- {
		}

		// Check scan complete
		if i >= j {
			break
		}

		a.Swap(i, j)
	}

	// Put partitioning item v into a[j]
	a.Swap(0, j)

	// whith a[lo..j-1 <= a[j] <= a[j+1..hi]
	return j
}

func median(a ds.ArrayList, l int) int {
	u := rand.Intn(l)
	v := rand.Intn(l)
	w := rand.Intn(l)

	x := a[u]
	y := a[v]
	z := a[w]

	if x < y {
		switch {
		case y < z:
			return v
		case x < z:
			return w
		default:
			return u
		}
	}
	switch {
	case x < z:
		return u
	case y < z:
		return w
	default:
		return v
	}
}

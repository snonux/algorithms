package sort

import (
	"algorithms/ds"
	"math/rand"
)

func Shuffle(a ds.ArrayList) ds.ArrayList {
	l := len(a)

	for i := 0; i < l; i++ {
		r := l - rand.Intn(l-i) - 1
		a.Swap(i, r)
	}

	return a
}

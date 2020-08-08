package sort

import (
	"algorithms/ds"
	"math/rand"
)

func Shuffle(a ds.ArrayList) ds.ArrayList {
	length := len(a)

	for i := 0; i < length; i++ {
		r := length - rand.Intn(length-i) - 1
		a.Swap(i, r)
	}

	return a
}

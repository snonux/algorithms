package ds

import (
	"fmt"
	"math/rand"
	"strings"
)

type ArrayList []int

func NewRandomArrayList(l, max int) ArrayList {
	a := make(ArrayList, l)
	for i := 0; i < l; i++ {
		if max > 0 {
			a[i] = rand.Intn(max)
			continue
		}
		a[i] = rand.Int()
	}
	return a
}

func NewAscendingArrayList(l int) ArrayList {
	a := make(ArrayList, l)
	for i := 0; i < l; i++ {
		a[i] = i
	}
	return a
}

func NewDescendingArrayList(l int) ArrayList {
	a := make(ArrayList, l)
	j := l - 1
	for i := 0; i < l; i++ {
		a[i] = j
		j--
	}
	return a
}

func (a ArrayList) FirstN(n int) string {
	var sb strings.Builder
	j := n

	l := len(a)
	if j > l {
		j = l
	}

	for i := 0; i < j; i++ {
		fmt.Fprintf(&sb, "%v ", a[i])
	}

	if j < l {
		fmt.Fprintf(&sb, "... ")
	}

	return sb.String()
}

func (a ArrayList) Sorted() bool {
	for i := len(a) - 1; i > 0; i-- {
		if a[i] < a[i-1] {
			return false
		}
	}
	return true
}

func (a ArrayList) Swap(i, j int) {
	aux := a[i]
	a[i] = a[j]
	a[j] = aux
}

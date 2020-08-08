package ds

import (
	"fmt"
	"math/rand"
	"strings"
)

type ArrayList []int

func NewRandomArrayList(length, max int) ArrayList {
	a := make(ArrayList, length)
	for i := 0; i < length; i++ {
		if max > 0 {
			a[i] = rand.Intn(max)
			continue
		}
		a[i] = rand.Int()
	}
	return a
}

func NewAscendingArrayList(length int) ArrayList {
	a := make(ArrayList, length)
	for i := 0; i < length; i++ {
		a[i] = i
	}
	return a
}

func NewDescendingArrayList(length int) ArrayList {
	a := make(ArrayList, length)
	j := length - 1
	for i := 0; i < length; i++ {
		a[i] = j
		j--
	}
	return a
}

func (a ArrayList) FirstN(n int) string {
	var sb strings.Builder
	j := n

	length := len(a)
	if j > length {
		j = length
	}

	for i := 0; i < j; i++ {
		fmt.Fprintf(&sb, "%v ", a[i])
	}

	if j < length {
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
	tmp := a[i]
	a[i] = a[j]
	a[j] = tmp
}

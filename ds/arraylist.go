package ds

import (
	"fmt"
	"strings"
)

type ArrayList []Integer

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
		if a[i].Lower(a[i-1]) {
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

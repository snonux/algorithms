package ds

import (
	"fmt"
	"math/rand"
)

type Integer struct {
	val int
}

func RandomIntegers(length, max int) []Comparer {
	a := make([]Comparer, length)
	for i := 0; i < length; i++ {
		a[i] = Integer{rand.Intn(max)}
	}
	return a
}

func SortedIntegers(length int) []Comparer {
	a := make([]Comparer, length)
	for i := 0; i < length; i++ {
		a[i] = Integer{i}
	}
	return a
}

func ReverseSortedIntegers(length int) []Comparer {
	a := make([]Comparer, length)
	j := length
	for i := 0; i < length; i++ {
		a[i] = Integer{j}
		j--
	}
	return a
}

func (i Integer) String() string {
	return fmt.Sprintf("%d", i.val)
}

func (i Integer) IntVal() int {
	return i.val
}

func (i Integer) LowerThan(j Comparer) bool {
	return i.val < j.IntVal()
}

func (i Integer) HigherThan(j Comparer) bool {
	return i.val > j.IntVal()
}

func (i Integer) Equals(j Comparer) bool {
	return i.val == j.IntVal()
}
